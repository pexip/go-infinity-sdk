#!/usr/bin/env sh

SCRIPT_PATH=$(realpath "$0")
SRC_DIR_PATH=$(dirname "$SCRIPT_PATH")
LIB_DIR_PATH="${SRC_DIR_PATH}/lib"

. "${LIB_DIR_PATH}/logging.sh"

main() {
  input_force="${1}"

  pip_installed=$(if command -v pip >/dev/null 2>&1; then echo true; else echo false; fi)
  echo "pip-installed=${pip_installed}" >> "$GITHUB_OUTPUT"

  bin_installed="false"
  if command -v zizmor >/dev/null 2>&1; then
    if [ "${input_force}" = "false" ]; then
      msg="Installation skipped."
      bin_installed="true"
    else
      msg="Executing forced installation."
    fi
    log_info "zizmor is found at $(which zizmor). ${msg}"
  else
    log_info "zizmor is not found. Executing installation."
  fi
  echo "bin-installed=${bin_installed}" >> "$GITHUB_OUTPUT"
}

main "$@"
