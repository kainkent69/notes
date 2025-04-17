## Valgrind Memcheck Usage Guide

This guide focuses on how to use Valgrind's Memcheck tool to detect memory errors in your C/C++ programs.

**Assumptions:**

* You have Valgrind already installed on your Arch Linux system.
* You are familiar with compiling C/C++ programs using `gcc` or `g++`.

**Key Concepts:**

* **Memcheck:** Valgrind's default tool. It detects memory leaks, invalid memory access (reads/writes), and other memory-related errors.
* **Debugging Information:** Compiler-generated information (using the `-g` flag) that allows Valgrind to provide precise error locations (file and line numbers).
* **Error Reports:** Memcheck's output, which describes the types of errors found and where they occurred.
* **Leak Summary:** A section in Memcheck's output that summarizes the memory leaks detected.

**1. Compilation with Debugging Info**

* Always compile your C/C++ code with the `-g` flag. This is essential for Valgrind to give you useful information.
* Example:
    * C:  `gcc -g -o myprog myprog.c`
    * C++: `g++ -g -o myprog myprog.cpp`

**2. Running Memcheck**

* To run your program with Memcheck, type `valgrind` followed by your program's execution command.
* Basic syntax:

    ```bash
    valgrind ./myprog
    ```

**3. Important Memcheck Options**

* `--leak-check=full`:  Provides detailed information about memory leaks.  *Always* use this for leak detection.
* `--show-reachable=yes`:  Shows memory blocks that are still pointed to but haven't been freed (reachable leaks).
* `-v`:  Verbose output.
* Common usage:

    ```bash
    valgrind --tool=memcheck --leak-check=full --show-reachable=yes ./myprog
    ```

**4. Interpreting Memcheck Output**

* Memcheck provides a lot of output, but focus on these key parts:
    * `ERROR SUMMARY`:  Indicates the total number of errors.  "0 errors from 0 contexts" is what you want.
    * Error reports:  For each error, Memcheck shows:
        * Error type (e.g., "Invalid write of size 4").
        * Memory address.
        * **Stack trace:** The sequence of function calls leading to the error.  This is *crucial* for finding the bug.  It includes file names and line numbers.
    * `LEAK SUMMARY`:  Shows how much memory was leaked.  With `--leak-check=full`, you'll see categories like "definitely lost".

**5. Examples**

* **Example 1: Memory Leak**
    * `leak.c`:

        ```c
        #include <stdlib.h>
        int main() {
            int *ptr = malloc(10 * sizeof(int));
            ptr[0] = 5;
            return 0; // Missing free(ptr);
        }
        ```

    * Compile: `gcc -g -o leak leak.c`
    * Run: `valgrind --tool=memcheck --leak-check=full ./leak`
    * Output will show a "definitely lost" leak in `main`.
* **Example 2: Invalid Write**
    * `invalid_write.c`:

        ```c
        #include <stdlib.h>
        int main() {
            int *arr = malloc(5 * sizeof(int));
            arr[5] = 10; // Writing beyond allocated size!
            free(arr);
            return 0;
        }
        ```

    * Compile: `gcc -g -o invalid_write invalid_write.c`
    * Run:  `valgrind ./invalid_write`
    * Output will report an "Invalid write of size 4" (assuming ints are 4 bytes).
* **Example 3: Use of Uninitialized Value**
    * `uninit.c`

        ```c
        #include <stdio.h>
        int main() {
           int x;
           printf("%d\n", x); // x is uninitialized
           return 0;
        }
        ```

    * Compile: `gcc -g -o uninit uninit.c`
    * Run: `valgrind ./uninit`
    * Output:  Will show "Conditional jump or move depends on uninitialised value(s)"

**6. Workflow Summary**

1.  Compile with `-g`.
2.  Run with `valgrind --tool=memcheck --leak-check=full ./your_program`.
3.  Carefully read the output.  Use the line numbers in the stack traces to find the errors in your source code.
4.  Fix your code, recompile, and re-run Valgrind until you get "0 errors" and "0 bytes in 0 blocks" leaked.

