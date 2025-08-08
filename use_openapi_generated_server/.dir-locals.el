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

                          ;; Profile 2: Attach to external debugger
                          (go-debug-docker
                           modes (go-mode go-ts-mode)
                           command "dlv"
                           port 40000
                           :request "attach"  ;; this will run "dlv attach ..."
                           :mode "remote"     ;; connect to a running debugger session
                           :type "go"
                           ;; https://github.com/emacs-lsp/dap-mode/blob/b97756665709bea37b9ffe262c5fa9196f1b4577/docs/page/configuration.md?plain=1#L622
                           ;; Set `from` as absolute path to repo root and map to code root in container.
                           :substitutePath (vector (ht ("from" (concat (funcall dape-cwd-fn) "use_openapi_generated_server")) ("to" "/src")))
                           )
                          )
                       )
         )
      )
 )
