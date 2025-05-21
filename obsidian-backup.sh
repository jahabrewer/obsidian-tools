#!/bin/bash

# Obsidian Backup Script

set -e # Exit immediately if a command exits with a non-zero status.

CONFIG_FILE="${HOME}/.obsidian-backup-config.yaml"
VERBOSE=false

# --- Helper Functions ---
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

verbose_log() {
    if [ "$VERBOSE" = true ]; then
        log "VERBOSE: $1"
    fi
}

# --- Check Dependencies ---
check_dependencies() {
    verbose_log "Checking dependencies..."
    if ! command -v yq &> /dev/null; then
        log "ERROR: yq is not installed. Please install yq to parse the config file."
        log "See: https://github.com/mikefarah/yq#install"
        exit 1
    fi
    verbose_log "yq found."

    if ! command -v zip &> /dev/null; then
        log "ERROR: zip is not installed. Please install zip."
        log "On macOS: brew install zip"
        log "On Debian/Ubuntu: sudo apt install zip"
        exit 1
    fi
    verbose_log "zip found."
    verbose_log "All dependencies satisfied."
}

# --- Read Configuration ---
read_config() {
    verbose_log "Reading configuration from ${CONFIG_FILE}..."
    if [ ! -f "$CONFIG_FILE" ]; then
        log "ERROR: Configuration file not found at ${CONFIG_FILE}"
        log "Please create it with your vault_path and backup_destination_path."
        log "Example content for ${CONFIG_FILE}:"
        log "vault_path: \"/path/to/your/obsidian_vault\""
        log "backup_destination_path: \"/path/to/your/backups_folder\""
        exit 1
    fi

    VAULT_PATH=$(yq e '.vault_path' "$CONFIG_FILE")
    BACKUP_DESTINATION_PATH=$(yq e '.backup_destination_path' "$CONFIG_FILE")

    if [ -z "$VAULT_PATH" ] || [ "$VAULT_PATH" == "null" ] || [ "$VAULT_PATH" == "/path/to/your/obsidian_vault" ]; then
        log "ERROR: vault_path is not set or is invalid in ${CONFIG_FILE}. Please update it."
        exit 1
    fi
    verbose_log "Vault path: ${VAULT_PATH}"

    if [ -z "$BACKUP_DESTINATION_PATH" ] || [ "$BACKUP_DESTINATION_PATH" == "null" ] || [ "$BACKUP_DESTINATION_PATH" == "/path/to/your/backups_folder" ]; then
        log "ERROR: backup_destination_path is not set or is invalid in ${CONFIG_FILE}. Please update it."
        exit 1
    fi
    verbose_log "Backup destination path: ${BACKUP_DESTINATION_PATH}"

    if [ ! -d "$VAULT_PATH" ]; then
        log "ERROR: Vault path '${VAULT_PATH}' does not exist or is not a directory."
        exit 1
    fi
}

# --- Perform Backup ---
perform_backup() {
    log "Starting Obsidian vault backup..."
    verbose_log "Source vault: ${VAULT_PATH}"
    verbose_log "Backup destination: ${BACKUP_DESTINATION_PATH}"

    # Determine zip options based on verbosity
    ZIP_OPTIONS="-r"
    if [ "$VERBOSE" = false ]; then
        ZIP_OPTIONS="-r -q"
    fi
    verbose_log "Using ZIP_OPTIONS: ${ZIP_OPTIONS}"

    # Ensure backup destination directory exists
    if ! mkdir -p "$BACKUP_DESTINATION_PATH"; then 
        log "ERROR: Could not create backup destination directory: ${BACKUP_DESTINATION_PATH}"
        exit 1
    fi
    verbose_log "Backup destination directory confirmed/created: ${BACKUP_DESTINATION_PATH}"

    TIMESTAMP=$(date +%Y-%m-%d_%H%M%S)
    VAULT_BASENAME=$(basename "${VAULT_PATH}")
    ZIP_FILENAME="${VAULT_BASENAME}-backup-${TIMESTAMP}.zip"
    TARGET_ZIP_PATH="${BACKUP_DESTINATION_PATH}/${ZIP_FILENAME}"

    log "Creating zip archive: ${TARGET_ZIP_PATH}"
    
    # Navigate to the parent directory of the vault to get clean paths in the zip
    # e.g., if VAULT_PATH is /Users/me/MyVault, cd to /Users/me, then zip MyVault
    VAULT_PARENT_DIR=$(dirname "${VAULT_PATH}")
    VAULT_DIR_NAME=$(basename "${VAULT_PATH}")

    # Check if we are trying to zip from root, which is problematic for cd
    if [ "$VAULT_PARENT_DIR" == "/" ] && [ "$VAULT_DIR_NAME" == "/" ]; then 
        # This case should ideally not happen if VAULT_PATH is a specific vault directory
        log "ERROR: Vault path seems to be the root directory. This is not supported for zipping in this manner."
        exit 1
    elif [ "$VAULT_PARENT_DIR" == "." ] && [ "$VAULT_DIR_NAME" == "$VAULT_PATH" ]; then
        # This happens if VAULT_PATH is a relative path like "MyVault" in the current dir
        # In this case, current directory is the parent for zip context
        verbose_log "Zipping relative path '${VAULT_DIR_NAME}' from current directory."
        if zip $ZIP_OPTIONS "${TARGET_ZIP_PATH}" "${VAULT_DIR_NAME}" -x "*.git/*" -x "*.DS_Store"; then
            log "SUCCESS: Backup created successfully at ${TARGET_ZIP_PATH}"
        else
            log "ERROR: Failed to create zip archive."
            # Attempt to remove partially created zip file if zip command failed
            [ -f "${TARGET_ZIP_PATH}" ] && rm "${TARGET_ZIP_PATH}"
            exit 1
        fi
    else
        verbose_log "Changing directory to ${VAULT_PARENT_DIR} for zipping."
        pushd "${VAULT_PARENT_DIR}" > /dev/null
        if zip $ZIP_OPTIONS "${TARGET_ZIP_PATH}" "${VAULT_DIR_NAME}" -x "*.git/*" -x "*.DS_Store"; then
            log "SUCCESS: Backup created successfully at ${TARGET_ZIP_PATH}"
        else
            log "ERROR: Failed to create zip archive."
            # Attempt to remove partially created zip file if zip command failed
            [ -f "${TARGET_ZIP_PATH}" ] && rm "${TARGET_ZIP_PATH}"
            popd > /dev/null
            exit 1
        fi
        popd > /dev/null
    fi
}

# --- Main Script Logic ---

# Parse command line arguments (e.g., -v for verbose, -c for different config)
while getopts "vc:" opt; do
  case $opt in
    v)
      VERBOSE=true
      ;;
    c)
      CONFIG_FILE=$OPTARG
      ;;
    \?)
      echo "Usage: cmd [-v] [-c <config_file>]"
      exit 1
      ;;
  esac
done

log "Obsidian Backup Script started."
if [ "$VERBOSE" = true ]; then log "Verbose mode enabled."; fi
if [ "$CONFIG_FILE" != "${HOME}/.obsidian-backup-config.yaml" ]; then log "Using custom config file: $CONFIG_FILE"; fi

check_dependencies
read_config
perform_backup

log "Obsidian Backup Script finished."

exit 0
