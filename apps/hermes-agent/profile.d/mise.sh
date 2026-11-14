# shellcheck shell=sh
for __mise_p in /opt/mise "${MISE_DATA_DIR:-${HOME}/.local/share/mise}/shims"; do
    case ":$PATH:" in
        *":$__mise_p:"*) ;;
        *) PATH="$__mise_p:$PATH" ;;
    esac
done
unset __mise_p
export PATH
