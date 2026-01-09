# Clipboard Translator

An intelligent real-time clipboard translation tool that automatically monitors your clipboard content and translates text to your desired language, seamlessly replacing the original text with the translation.

## 🎯 Overview

This project provides a hands-free translation experience directly from your clipboard:
- **Real-time Monitoring**: Continuously watches clipboard for new content
- **Automatic Translation**: Instant translation without manual intervention
- **Language Support**: Configurable target language for translation
- **Seamless Integration**: Works with any application that uses clipboard
- **API Integration**: Leverages external translation services for accuracy
- **Background Operation**: Runs silently in the background

## 🛠️ Technology Stack

- **Language**: Go 1.21+
- **Clipboard Access**: `atotto/clipboard` for cross-platform clipboard operations
- **Configuration**: `godotenv` for secure API key management
- **HTTP Client**: Native Go HTTP client for API communication
- **Environment Variables**: Secure credential storage

## 🚀 Features

### Core Functionality
- **Continuous Monitoring**: Real-time clipboard change detection
- **Automatic Translation**: Instant translation when new content is detected
- **Multi-language Support**: Configurable target languages (en, uk, es, etc.)
- **API Integration**: External translation service integration
- **Error Handling**: Robust error handling with graceful degradation
- **Memory Management**: Efficient clipboard content tracking

### Translation Workflow
1. Monitor clipboard for changes
2. Detect new text content
3. Send text to translation API
4. Receive translated text
5. Replace clipboard content with translation
6. Continue monitoring for next change

### Supported Languages
- English (en)
- Ukrainian (uk)
- Spanish (es)
- And more depending on translation API capabilities

## 📋 Prerequisites

- Go 1.21 or higher installed
- Translation API access (API key required)
- Git for cloning the repository
- Valid translation service credentials

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd translator
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   Create a `.env` file in the project root:
   ```env
   API_KEY=your_translation_api_key
   URL=your_translation_service_endpoint
   ```

4. **Build the application**
   ```bash
   go build -o clipboard-translator main.go
   ```

## 🎮 Usage

### Running the Application

```bash
# Using go run
go run main.go

# Using built binary
./clipboard-translator
```

### Setup Process

1. **Start the application**
   ```bash
   go run main.go
   ```

2. **Select target language**
   ```
   Input target language like: en, uk, es: en
   ```

3. **Start translating**
   - Copy any text to your clipboard (Ctrl+C or Cmd+C)
   - The application will automatically detect the change
   - Text will be translated and replaced in your clipboard
   - Paste the translated text anywhere (Ctrl+V or Cmd+V)

### Example Workflow

```bash
# Start the translator
$ go run main.go
Input target language like: en, uk, es: en

# In another application, copy text: "Bonjour le monde"
# The application detects this and translates it
# Paste anywhere: "Hello world"

# Copy more text: "¿Cómo estás?"
# Automatically translated and ready to paste: "How are you?"
```

## 🏗️ Project Structure

```
translator/
├── main.go              # Main application entry point
├── .env                 # Environment variables (create this)
├── go.mod               # Go module file
├── go.sum               # Go module checksums
└── README.md            # This file
```

## 🧩 Core Components

### Configuration Structure
```go
type EnvData struct {
    apiKey string
    url    string
}
```

### Key Systems

1. **Clipboard Monitoring**
   - Continuous polling for clipboard changes
   - Content comparison to detect new text
   - Memory-efficient change detection

2. **Translation Service**
   - HTTP client for API communication
   - Request/response handling
   - Error handling and retry logic

3. **Language Management**
   - Target language configuration
   - Language code validation
   - API parameter formatting

4. **Background Processing**
   - Infinite loop for continuous operation
   - Graceful error handling
   - Resource cleanup

## 🔧 Configuration

### Environment Variables

| Variable | Description | Required | Example |
|----------|-------------|----------|---------|
| `API_KEY` | Translation service API key | Yes | `your_api_key_here` |
| `URL` | Translation service endpoint | Yes | `https://api.translate.service/v1/translate` |

### Language Codes

Commonly supported language codes:
- `en` - English
- `uk` - Ukrainian
- `es` - Spanish
- `fr` - French
- `de` - German
- `it` - Italian
- `pt` - Portuguese
- `ru` - Russian
- `ja` - Japanese
- `zh` - Chinese

### Translation API Format

The application expects the translation API to accept:
- Source text as input
- Target language parameter
- API key for authentication
- Return translated text as response

## 🎯 Performance Considerations

- **Efficient Monitoring**: Minimal CPU usage with smart change detection
- **Memory Management**: Low memory footprint with efficient string handling
- **Network Optimization**: Single HTTP request per translation
- **Error Recovery**: Graceful handling of network failures

## 🧪 Integration Examples

### Use Cases

1. **Document Translation**
   - Copy text from documents
   - Automatically translate to target language
   - Paste translated content

2. **Web Browsing**
   - Copy foreign language content
   - Instant translation for understanding
   - Seamless multilingual browsing

3. **Communication**
   - Translate messages in real-time
   - Cross-language communication
   - Social media content translation

4. **Development**
   - Translate error messages
   - Understand foreign code comments
   - Internationalization testing

### Application Compatibility

Works with any application that uses system clipboard:
- Web browsers
- Text editors
- Office applications
- Messaging apps
- Development tools
- Terminal emulators

## 🚧 Development Notes

### Extending the Application
- Add support for multiple translation services
- Implement language auto-detection
- Add translation history logging
- Create GUI interface for better UX
- Add batch translation capabilities
- Implement translation caching

### Potential Improvements
- Add translation quality indicators
- Implement custom dictionary support
- Add voice output for translations
- Create browser extension integration
- Add translation confidence scoring
- Implement offline translation capabilities

### Known Limitations
- Requires internet connection
- Dependent on external API availability
- No translation history
- No batch processing
- Limited to text content
- API rate limits may apply

## 🔒 Security Considerations

- **API Key Protection**: Store API keys in environment variables
- **Data Privacy**: Clipboard content is sent to external service
- **Network Security**: Ensure HTTPS endpoints for translation API
- **Access Control**: Limit API key permissions as needed

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## 📞 Contact

For questions or support regarding this translator, please open an issue in the repository.

---

**Built with ❤️ using Go**
