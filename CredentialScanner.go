 func checkSecurityPatterns(line string, path string) {
    // ... önceki kontroller ...

    // 1. Credential (API Key/Token) Sızıntısı - (Informational/Low-Medium)
    // VS Code'un veya eklentilerin gizli token'ları hardcoded bırakıp bırakmadığını kontrol eder.
    if strings.Contains(line, "api_key") || strings.Contains(line, "secret") || strings.Contains(line, "password") {
        if strings.Contains(line, "=") {
             addFinding("CREDENTIAL_LEAK_RISK", path, line)
        }
    }

    // 2. Güvensiz API Kullanımı (Erişim Kontrolü)
    // VS Code'un `Context` veya `Configuration` API'lerini kullanan eklentilerin,
    // yetkisiz okuma yapıp yapmadığını anlamak için bu noktalar önemlidir.
    if strings.Contains(line, "workspace.getConfiguration") {
        addFinding("CONFIG_ACCESS_POINT", path, line)
    }

    // 3. Command Injection (SQLi/Command Injection'a benzer yapılar)
    // Eğer VS Code terminale bir komut basıyorsa ve bu komut `template string` (${...}) içeriyorsa çok tehlikelidir.
    if strings.Contains(line, "terminal.sendText") && strings.Contains(line, "${") {
        addFinding("COMMAND_INJECTION_RISK", path, line)
    }
}
