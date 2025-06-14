
# Development Design Document for LumberJack Monitoring System
Testing...
## Overview
The LumberJack Monitoring System is a full-stack application designed to monitor system resources and provide real-time analytics. It consists of a **Go backend** for API services and a **React frontend** for the user interface. The application dynamically fetches and displays data such as CPU usage, network traffic, user information, and storage analytics.

---

## Repository Layout

### Root Directory
```
├── server/         # Backend codebase (Go)
├── client/         # Frontend codebase (React + TypeScript)
```

---

## Backend (`server/`)

The backend is written in **Go** and provides RESTful APIs for the frontend to fetch system data. It is structured into modular packages for separation of concerns.

### Directory Structure
```
server/
├── cmd/            # Main entry point for the server
│   └── system/
│       ├── sys-scan.go         # Main server logic
│       ├── sys-scan_test.go    # Unit tests for the server
├── pkg/            # Core packages for functionality
│   ├── cpu/                   # CPU-related utilities
│   ├── logging/               # Error logging utilities
│   ├── network/               # Network-related utilities
│   ├── remoteConnection/      # Remote connection utilities
│   ├── routes/                # API route definitions
│   └── userSpace/             # User-related utilities
├── go.mod          # Go module definition
```

### Key Components

#### 1. **Main Server (`cmd/system/sys-scan.go`)**
- **Purpose**: Entry point for the backend server.
- **Functionality**:
  - Registers API routes using `routes.RegisterRoutes()`.
  - Starts the HTTP server on port `8080`.
  - Logs errors using the `logging` package.
- **Example**:
  ```go
  routes.RegisterRoutes()
  http.ListenAndServe(":8080", nil)
  ```

#### 2. **Routes (`pkg/routes/routes.go`)**
- **Purpose**: Defines all API endpoints.
- **Endpoints**:
  - `/api/home`: Fetches a welcome message, CPU usage, user IP, and user info.
  - `/api/cpu`: Returns CPU usage data.
  - `/api/network`: Placeholder for network usage data.
  - `/api/storage`: Placeholder for storage analytics.
  - `/api/user`: Fetches user and server information.
- **Dynamic Data**: Data is fetched dynamically using utility functions from other packages (e.g., `cpu.GetUsage()` and `userSpace.GetUserData()`).

#### 3. **User Utilities (`pkg/userSpace/userSpace.go`)**
- **Purpose**: Provides user-related data.
- **Functions**:
  - `GetUsername()`: Returns the current user's username.
  - `GetUserData()`: Returns detailed user information (e.g., username, user ID, home directory, OS).
- **Example**:
  ```go
  userData, err := userSpace.GetUserData()
  if err != nil {
      fmt.Println("Error fetching user data:", err)
  }
  ```

#### 4. **CPU Utilities (`pkg/cpu/cpu-scan.go`)**
- **Purpose**: Provides CPU-related data.
- **Functions**:
  - `GetUsage()`: Returns the number of CPU threads.
- **Example**:
  ```go
  cpuInfo := cpu.GetUsage()
  ```

#### 5. **Network Utilities (`pkg/network/network-traffic.go`)**
- **Purpose**: Handles network-related operations.
- **Functions**:
  - `ReadUserIP()`: Extracts the user's IP address from the HTTP request headers.
  - `GetNetData()`: Placeholder for network traffic data.
- **Example**:
  ```go
  userIP := network.ReadUserIP(r)
  ```

#### 6. **Remote Connection Utilities (`pkg/remoteConnection/rc.go`)**
- **Purpose**: Handles remote connection data.
- **Functions**:
  - `GetIP()`: Extracts the IP address from HTTP headers or `RemoteAddr`.
  - `SSHAttemptHandler()`: Placeholder for handling SSH attempts.

#### 7. **Logging (`pkg/logging/errorLog.go`)**
- **Purpose**: Provides error logging functionality.
- **Functions**:
  - `WriteLog(err error)`: Logs errors with timestamps and file/line information.
  - `WriteLogWithMsg(msg string)`: Logs custom messages.
- **Example**:
  ```go
  logging.NewErrorLog().WriteLog(err)
  ```

---

## Frontend (`client/`)

The frontend is built with **React**, **TypeScript**, and **Vite**. It provides a dynamic and responsive user interface for monitoring system data.

### Directory Structure
```
client/
├── lumber-jack/
│   ├── src/                # Source code
│   │   ├── App.tsx         # Main application component
│   │   ├── main.tsx        # Entry point for React
│   │   ├── App.css         # Component-specific styles
│   │   ├── index.css       # Global styles
│   │   └── vite-env.d.ts   # Vite environment types
│   ├── vite.config.ts      # Vite configuration
│   ├── tsconfig.json       # TypeScript configuration
│   ├── package.json        # NPM dependencies and scripts
│   ├── .gitignore          # Git ignore rules
│   └── README.md           # Documentation
```

### Key Components

#### 1. **Main Application (`src/App.tsx`)**
- **Purpose**: Core React component that manages the UI and fetches data from the backend.
- **Features**:
  - **Dynamic Tabs**: Users can switch between tabs (e.g., Home, Network, CPU, Storage, User Info).
  - **Data Fetching**: Fetches data dynamically from the backend using the `fetch` API.
  - **Loading State**: Displays a loading indicator while data is being fetched.
- **Example**:
  ```tsx
  const fetchData = (tab: string) => {
    setLoading(true);
    fetch(`/api/${tab}`)
      .then((res) => res.json())
      .then((fetchedData) => {
        setData(fetchedData);
        setLoading(false);
      });
  };
  ```

#### 2. **Styling**
- **Global Styles (`src/index.css`)**:
  - Defines global styles such as font, colors, and layout.
  - Supports light and dark themes using `prefers-color-scheme`.
- **Component Styles (`src/App.css`)**:
  - Styles for individual components like data cards and the sidebar.
  - Responsive design for smaller screens.

#### 3. **Routing**
- The frontend uses a tab-based navigation system. Each tab corresponds to a backend API endpoint (e.g., `/api/home`, `/api/cpu`).

#### 4. **Vite Configuration (`vite.config.ts`)**
- **Proxy Setup**: Proxies API requests to the Go backend running on `http://localhost:8080`.
- **Plugins**: Uses `@vitejs/plugin-react` for React support.

---

## API Endpoints

| Endpoint       | Description                          | Backend Functionality                     |
|----------------|--------------------------------------|-------------------------------------------|
| `/api/home`    | Fetches welcome message and system data. | Combines CPU, user, and IP data.          |
| `/api/cpu`     | Fetches CPU usage data.              | Returns the number of CPU threads.        |
| `/api/network` | Placeholder for network usage data.  | Returns static or dynamic data.           |
| `/api/storage` | Placeholder for storage analytics.   | Returns static or dynamic data.           |
| `/api/user`    | Fetches user and server information. | Returns user details like username and OS.|

---

## Development Practices

1. **Backend**:
   - Modular design with separate packages for CPU, network, user, and logging functionalities.
   - Error handling with descriptive messages and logging.
   - Dynamic data fetching using Go's standard library.

2. **Frontend**:
   - Component-based architecture with React.
   - TypeScript for type safety.
   - Responsive design for better user experience.

3. **Testing**:
   - Unit tests for backend routes (e.g., `sys-scan_test.go`).
   - Mocked HTTP requests for testing API responses.

---

## Future Improvements
1. **Dynamic Data**:
   - Implement real-time network and storage analytics.
   - Replace placeholders with actual data.

2. **Error Handling**:
   - Improve error messages in the frontend for better debugging.

3. **Authentication**:
   - Add user authentication for secure access.

4. **Frontend Enhancements**:
   - Use a state management library (e.g., Redux) for better state handling.
   - Add routing for better navigation.

---

This document provides a comprehensive overview of the LumberJack Monitoring System's current implementation and serves as a guide for further development.
```
