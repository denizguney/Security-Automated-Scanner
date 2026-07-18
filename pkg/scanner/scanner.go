func checkSecurityPatterns(line string, path string) {
    // 1. RCE Riskleri (Child Process)
    if strings.Contains(line, "child_process.exec") || strings.Contains(line, "spawn(") {
        fmt.Printf("[CRITICAL] RCE Riski: %s\n", path)
    }

    // 2. WebView XSS Riskleri
    if strings.Contains(line, ".innerHTML") || strings.Contains(line, "webview.html =") {
        fmt.Printf("[HIGH] Potansiyel XSS (WebView): %s\n", path)
    }

    // 3. IPC Haberleşme (Potansiyel Yetki Yükseltme / RCE)
    if strings.Contains(line, "ipc.on") || strings.Contains(line, "ipc.send") {
        fmt.Printf("[MEDIUM] IPC Mesajlaşma Noktası: %s\n", path)
    }
    
    // 4. Güvensiz Dosya İşlemleri (Path Traversal)
    if strings.Contains(line, "fs.readFile") || strings.Contains(line, "fs.writeFile") {
        if strings.Contains(line, "join(") {
            fmt.Printf("[MEDIUM] Potansiyel Path Traversal: %s\n", path)
        }
    }
}
