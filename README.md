## GopenAI
A lightweight Go library for interacting with the OpenAI API. It simplifies the process of making API requests and handling responses. 

### Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [(Planned) Features](#planned-features)
- [Contributing](#contributing)
- [License](#license)

### Installation
```bash
go get github.com/eloonstra/gopenai
```

### Usage
1. Import the package.
```go
import "github.com/eloonstra/gopenai"
```
2. Create a new client.
```go
ai := gopenai.New("YOUR_API_KEY")
```
3. Build a request.
```go
request := ai.CreateCompletion("MODEL_ID").
    WithPrompt("This is a test prompt").
    WithMaxTokens(10).
    WithTemperature(0.5)
```

4. After building your request, you can call the `Do()` method to execute the request.
```go
response, err := request.Do()
```

### (Planned) Features
A list of features that have already been implemented and those that are planned.
- [x] Chat
- [x] Completion
- [x] Edit
- [x] Model
- [x] Moderation
- [ ] Chat (stream)
- [ ] Completion (stream)
- [ ] File
- [ ] Audio
- [ ] Image
- [ ] Embedding
- [ ] Fine-tune

### Contributing
Contributions are welcome! Please open an issue or pull request if you would like to contribute.

### License
GopenAI is licensed under the MIT License. See [LICENSE](LICENSE) for more information.