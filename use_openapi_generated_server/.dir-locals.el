(
 (nil . (
         (dape-configs . (
                          ;; https://dev.to/dorneanu/mastering-golang-debugging-in-emacs-34p7
                          ;; Profile 1: Launch application and start DAP server
                          (go-debug-locally
                           modes (go-mode go-ts-mode)
                           command "dlv"
                           ;; Debug port != Application port.
                           command-args ("dap" "--listen" "127.0.0.1:5678")
                           command-cwd default-directory
                           host "127.0.0.1"
                           port 5678
                           :request "launch"
                           :mode "debug"
                           :type "go"
                           :showLog "true"
                           :program default-directory)
                          ;; TODO: Add Tramp config to attach to remote go app.
                          ;; Also need an SSH tunnel! See:
                          ;; https://github.com/svaante/dape/issues/109
                          )
                       )
         )
      )
 )
