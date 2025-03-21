# Modern Todo App with HTMX and Go

A lightweight, fast Todo application built with Go and HTMX, featuring a clean dark theme UI and real-time updates.

## ✨ Features

- ⚡ Real-time updates with HTMX
- 🎨 Modern dark theme UI with TailwindCSS
- 🚀 Fast and lightweight
- 💾 PostgreSQL database storage
- 🔄 No page refreshes needed

## 🛠️ Tech Stack

- Go (Backend)
- HTMX (Frontend interactions)
- TailwindCSS & DaisyUI (Styling)
- PostgreSQL (Database)
- Gorilla Mux (Router)

## 📋 Prerequisites

- Go 1.16 or higher
- PostgreSQL
- Node.js (for TailwindCSS)

## 🚀 Quick Start

1. Clone the repository:

   ```bash
   git clone <your-repo-url>
   cd todo-htmx
    ```

2. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
    ```

3. Create PostgreSQL database:
   ```bash
   createdb todos
    ```

4. Run the application:
   ```bash
   go run main.go
    ```

The app will be available at localhost:5000

## 📝 API Endpoints

- GET / - Home page
- POST /todo - Create new todo
- PUT /todo/{id} - Mark todo as done
- DELETE /todo/{id} - Delete todo

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.
