# 🌸 Weather App

A beautiful, multilingual weather application built with Go and modern web technologies. Features interactive maps, voice commands, and cloud coverage visualization.

## ✨ Features

### 🌍 Multilingual Support
- **English** 🇺🇸
- **German** 🇩🇪  
- **Ukrainian** 🇺🇦
- **Hebrew** 🇮🇱 (with RTL support)

### 🗺️ Interactive Map
- **OpenStreetMap** integration with Leaflet
- **Click to select** any location worldwide
- **Current location** detection with permission
- **Custom pink-themed** markers and controls
- **Dynamic cloud coverage** based on weather conditions

### 🎤 Voice Interface
- **Speech recognition** for hands-free operation
- **Text-to-speech** weather announcements
- **Multilingual voice commands**
- **Voice reading toggle** with persistence

### 🌤️ Weather Features
- **Real-time weather data** from OpenWeatherMap API
- **Temperature, humidity, wind speed**
- **Weather descriptions** in multiple languages
- **Cloud coverage visualization**
- **Animated cloud layer** for atmosphere

### 🎨 Beautiful UI
- **Pink-rose theme** with 3D effects
- **Responsive design** for all devices
- **Smooth animations** and transitions
- **Glass morphism** effects
- **Floating cloud animations**

## 🚀 Quick Start

### Prerequisites
- Go 1.19 or higher
- OpenWeatherMap API key

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd weather-go
   ```

2. **Set your API key**
   ```bash
   export OPENWEATHERMAP_API_KEY=your_api_key_here
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

4. **Open your browser**
   Navigate to `http://localhost:58080`

## 🎤 Voice Commands

### English
- "weather in [city]"
- "temperature in [city]"
- "current weather"
- "weather here"
- "help"

### German
- "wetter in [stadt]"
- "temperatur in [stadt]"
- "aktuelles wetter"
- "wetter hier"
- "hilfe"

### Ukrainian
- "погода в [місто]"
- "температура в [місто]"
- "поточна погода"
- "погода тут"
- "допомога"

### Hebrew
- "מזג אוויר ב [עיר]"
- "טמפרטורה ב [עיר]"
- "מזג אוויר נוכחי"
- "מזג אוויר כאן"
- "עזרה"

## 🛠️ Configuration

### Environment Variables
- `OPENWEATHERMAP_API_KEY`: Your OpenWeatherMap API key

### Local Storage
The app automatically saves:
- **Language preference**
- **Voice reading toggle state**

## 🏗️ Project Structure

```
weather-go/
├── main.go              # Go server entry point
├── templates/
│   └── index.html       # Main web interface
├── go.mod               # Go module file
├── go.sum               # Go dependencies
├── .gitignore           # Git ignore rules
└── README.md            # This file
```

## 🔧 Technical Stack

- **Backend**: Go with Gin framework
- **Frontend**: HTML5, CSS3, JavaScript (ES6+)
- **Maps**: OpenStreetMap with Leaflet.js
- **Weather API**: OpenWeatherMap
- **Voice**: Web Speech API
- **Storage**: Browser localStorage

## 🌟 Key Features

### Responsive Design
- Works on desktop, tablet, and mobile devices
- Adaptive layout for different screen sizes
- Touch-friendly interface

### Accessibility
- Screen reader compatible
- Keyboard navigation support
- High contrast elements
- Voice command alternatives

### Performance
- Optimized cloud animations
- Efficient map rendering
- Minimal API calls
- Fast page loading

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🙏 Acknowledgments

- **OpenWeatherMap** for weather data
- **OpenStreetMap** for map tiles
- **Leaflet.js** for map functionality
- **Web Speech API** for voice features

---

**Enjoy your weather experience! 🌸☁️** 