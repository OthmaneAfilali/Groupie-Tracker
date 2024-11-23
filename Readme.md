# Groupie Tracker

## Overview

Groupie Tracker is a dynamic and interactive web application that showcases information about bands and artists using a provided API. The website provides detailed data visualizations, including information about artists, their members, concert dates, and locations. It supports server-client interaction to deliver a seamless and engaging user experience.

This project is built using **Go** and adheres to robust development practices, ensuring smooth operation and error handling.

---

## Features

### Core Functionalities:
- **Artist Profiles**: View artist details, including name, image, members, year of activity, and first album date.
- **Concert Info**: Explore past and upcoming concert locations and dates.
- **Data Visualizations**: User-friendly presentation of data through cards, tables, lists, and graphics.
- **Event-Driven Actions**: Features dynamic client-server interaction, such as submitting artist names via forms and retrieving specific data in real time.
  
### Highlights:
- **Responsive Design**: Styled for optimal viewing across devices.
- **Error Handling**: Displays informative error pages to ensure smooth navigation.
- **Testable Codebase**: Includes unit tests for verifying application functionality.

---

## Project Structure

### Files and Folders:
```plaintext
/groupie-tracker
│
├── main.go
├── handlers.go
├── api.go
├── utils.go
├── templates/
│   ├── index.html
│   ├── about.html
│   ├── bio.html
│   └── error.html
├── assets/
│   ├── style.css
│   └── script.js
```
## File Descriptions

### `main.go`
The application's entry point:
- Initializes and validates data from the API.
- Configures HTTP routes.
- Launches the web server.

### `handlers.go`
Manages server-side logic for different routes:
- **`homeHandler`**: Renders the home page.
- **`aboutHandler`**: Displays the about page.
- **`bioHandler`**: Handles POST and GET requests for artist details.

### `api.go`
Handles API communication and data parsing:
- **`fetchAllData`**: Aggregates all API data and prepares it for rendering.
- **`fetchAndUnmarshal`**: Fetches JSON data and converts it to Go structures.
- **`FetchApi`**: Makes HTTP GET requests.
- **`validateData`**: Ensures data correctness and completeness.

### `utils.go`
Utility functions for error management:
- **`errorHandler`**: Renders error pages.
- **`ErrorCheck`**: Identifies and logs runtime errors.
- **`logError`**: Custom error logger for debugging.

### `templates/`
Contains HTML files for rendering:
- `index.html`: Home page template.
- `about.html`: About the application.
- `bio.html`: Artist details.
- `error.html`: Error messages and handling.

### `assets/`
Static files for enhancing the user experience:
- `style.css`: Custom styles for the application.
- `script.js`: Client-side logic, including form handling and dynamic interactions.

---

## Setup and Usage

### Prerequisites
- **Go**: Version 1.19 or higher.
- A compatible web browser.
- An active internet connection for API data.

### Running the Application
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd groupie-tracker
2. Execute the program:
   ```bash
   go run main.go
3. Open your browser and visit:
    ```bash
    http://localhost:8080/groupie-tracker

## API Integration

This application integrates with the [Groupie Tracker API](https://groupietrackers.herokuapp.com/api). The API is divided into four endpoints:

1. **`artists`**: Provides general information about bands and artists.
2. **`locations`**: Contains concert locations.
3. **`dates`**: Provides concert schedules.
4. **`relation`**: Links all the above data for a cohesive view.

---

## Contributing

Contributions are welcome! If you'd like to contribute:
1. Fork the repository.
2. Create a new branch (`feature/my-feature`).
3. Submit a pull request with detailed changes.

Please follow good coding practices and ensure your contributions include relevant tests.
