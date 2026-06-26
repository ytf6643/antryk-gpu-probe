package main

import (
      "encoding/json"
      "net/http"
      "os"
      "os/exec"
  )

func main() {
      http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                b, e := exec.Command("nvidia-smi").CombinedOutput()
                s := string(b)
                if e != nil {
                              s = e.Error() + "\n" + s
                          }
                json.NewEncoder(w).Encode(map[string]string{
                              "status":               "ok",
                              "cuda_visible_devices": os.Getenv("CUDA_VISIBLE_DEVICES"),
                              "nvidia_smi":           s,
                          })
            })
      http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
                json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
            })
      p := os.Getenv("PORT")
      if p == "" {
                p = "8000"
            }
      http.ListenAndServe(":"+p, nil)
  }
