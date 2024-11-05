 - 📝 To do 
 - ✅ Done

### **Phase 1: Basics of Go**

#### **1. Go Installation & Setup** ✅

-   **Learn**: How to install Golang and set up your Go workspace.
-   **Action**: Install Go, configure your environment, and set up your IDE.
-   **Output**: A basic “Hello, World!” program.

#### **2. Basic Syntax and Control Structures** ✅

-   **Learn**: Variables, loops, conditionals, and functions in Go.
-   **Action**: Write simple programs to practice these concepts.
-   **Project Task**: Create a simple command-line application that outputs static information about the system (e.g., mock CPU, memory usage).

----------

### **Phase 2: Building the Core CLI Features**

#### **3. CLI Basics & Command-Line Arguments** ✅

-   **Learn**: How to handle command-line arguments and create flags using `flag` package.
-   **Action**: Add flags to select specific metrics (CPU, Memory, Storage).
-   **Project Task**: Implement a CLI structure to parse arguments like `--cpu`, `--memory`, `--network`, and `--storage`.

#### **4. Structs & Functions** ✅

-   **Learn**: How to define and use structs, methods, and interfaces.
-   **Action**: Organize your system metrics into different structs (e.g., `CPUInfo`, `MemoryInfo`).
-   **Project Task**: Create a function for each metric (CPU, memory, storage) that returns static information.

----------

### **Phase 3: Fetching System Information**

#### **5. Libraries and Packages** ✅

-   **Learn**: How to import and use third-party packages in Go.
-   **Action**: Explore libraries like `gopsutil` for fetching system metrics.
-   **Project Task**: Implement dynamic CPU and memory fetching using `gopsutil`.

#### **6. Error Handling** ✅

-   **Learn**: Proper error handling in Go using `error` type and handling functions.
-   **Action**: Add error handling for cases where fetching system information fails.
-   **Project Task**: Ensure that your CLI gracefully handles errors and prints useful messages.

----------

### **Phase 4: Improving Functionality**

#### **7. File I/O and Logging** ✅

-   **Learn**: How to handle files and logging in Go.
-   **Action**: Allow the CLI to log system information into a file.
-   **Project Task**: Add a `--log` flag that saves the system state output to a file.

#### **8. Goroutines & Concurrency** ✅

-   **Learn**: Basic concurrency with goroutines and channels in Go.
-   **Action**: Use goroutines to fetch different metrics concurrently.
-   **Project Task**: Optimize your CLI by fetching CPU, memory, and storage info simultaneously using goroutines.

----------

### **Phase 5: Real-Time Monitoring and Enhancements**

#### **9. Timers and Intervals** ✅

-   **Learn**: How to use `time` package for scheduling tasks.
-   **Action**: Implement real-time monitoring with intervals using Go’s ticker functionality.
-   **Project Task**: Add a `--monitor` flag that refreshes the system state every few seconds.

#### **10. Formatting and Output** 📝

-   **Learn**: Output formatting techniques (JSON, tables) and customizing the CLI output.
-   **Action**: Present system information in a more readable format (e.g., table format or JSON output).
-   **Project Task**: Add support for outputting system metrics in JSON or human-readable table format.

----------

### **Phase 6: Project Cleanup and Future Enhancements**

#### **11. Refactoring and Code Organization** ✅

-   **Learn**: How to organize large Go projects into packages and modules.
-   **Action**: Split your project into logical packages (e.g., `cpu`, `memory`, `storage`).
-   **Project Task**: Refactor your code into modular packages for better maintainability.

#### **12. Testing and Validation** ✅

-   **Learn**: Unit testing in Go with `testing` package.
-   **Action**: Write tests for the key functions of your CLI.
-   **Project Task**: Add basic unit tests for system info fetching and argument parsing.

#### **13. Deployment & Documentation** ✅

-   **Learn**: How to build, deploy, and distribute your CLI tool.
-   **Action**: Build cross-platform binaries and write comprehensive documentation.
-   **Project Task**: Prepare your project for open-source contribution by documenting the setup process, usage, and contribution guidelines.