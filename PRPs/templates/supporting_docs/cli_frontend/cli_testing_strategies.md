##CLI Testing Strategies

### 1. Command Testing
```
// cli/internal/commands/agentical_test.go
func TestAgenticalCommand(t *testing.T) {
    // Test command parsing
    cmd := &cobra.Command{}
    args := []string{"--config", "test.yaml", "run"}
    
    err := cmd.Execute()
    assert.NoError(t, err)
}

// CLI output testing
func TestAgenticalOutput(t *testing.T) {
    var buf bytes.Buffer
    cmd := NewAgenticalCommand()
    cmd.SetOut(&buf)
    cmd.SetArgs([]string{"status"})
    
    err := cmd.Execute()
    assert.NoError(t, err)
    assert.Contains(t, buf.String(), "Agentical Status:")
}
```

### 2. Integration Testing with Real CLI
```
// cli/integration/cli_integration_test.go
func TestCLIFullWorkflow(t *testing.T) {
    // Build CLI binary for testing
    cmd := exec.Command("go", "build", "-o", "testcli", "./cmd/devqai")
    err := cmd.Run()
    require.NoError(t, err)
    defer os.Remove("testcli")
    
    // Test actual CLI execution
    output, err := exec.Command("./testcli", "agentical", "list").Output()
    require.NoError(t, err)
    assert.Contains(t, string(output), "Available agents:")
}
```

### 3. Testing Charm TUI Components
```
// Use Bubble Tea's teatest for TUI testing
func TestInputModel(t *testing.T) {
    tm := teatest.NewTestModel(t, initialModel(), teatest.WithInitialTermSize(300, 100))
    
    // Send keystrokes
    tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("test input")})
    tm.Send(tea.KeyMsg{Type: tea.KeyEnter})
    
    // Assert final output
    teatest.RequireEqualOutput(t, goldenFile)
}
```
## trustlis Testing

### 1. Cryptographic Function Testing
```
// trustlis/crypto_test.go
func TestZKSTARKGeneration(t *testing.T) {
    // Test deterministic crypto operations
    input := []byte("test data")
    proof1 := GenerateZKSTARK(input)
    proof2 := GenerateZKSTARK(input)
    
    assert.True(t, VerifyZKSTARK(proof1, input))
    assert.Equal(t, proof1, proof2) // Should be deterministic
}

func TestAuthTokenGeneration(t *testing.T) {
    userID := "test-user"
    token := GenerateAuthToken(userID)
    
    claims, err := ValidateAuthToken(token)
    require.NoError(t, err)
    assert.Equal(t, userID, claims.UserID)
}
```

### 2. Auth Flow Integration Testing
```
// trustlis/integration/auth_flow_test.go
func TestCompleteAuthFlow(t *testing.T) {
    // Start trustlis server
    server := startTestServer(t)
    defer server.Close()
    
    // Test registration
    user := registerTestUser(t, server.URL)
    
    // Test authentication
    token := authenticateUser(t, server.URL, user)
    
    // Test API access with token
    response := makeAuthenticatedRequest(t, server.URL, token)
    assert.Equal(t, http.StatusOK, response.StatusCode)
}
```
