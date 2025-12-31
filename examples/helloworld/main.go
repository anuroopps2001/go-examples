// package = project or workspace
/* Package is a collection of common source code files

   Package can each related file ending with .go such as main.go, support.go and helper.go etc

   If all files belong to same package, then on each file at the beginning they should have:

   package <package_name>

   Types of packages
   1. Executable packages: which are used to do something, like when we run go build, it generates an executable which we use to run the program

   2. Reusuable packages: Defines/Generates a package that can be used as dependency  Ex: package apple


Name of the packages inside the source code defines, whether its creating executable or reusuable packages
Package with name "main" , will generate executable packages and that's how main.exe was created when we run go build main.go in this case.const




*/
package main

/* Import statement is used to give our package i,e "main", access the some code written inside another package i.e "fmt" in this case.

fmt library is used to print lot of things on to the terminal and we are making use of fmt package inside our main package

Basically, to use other packages along with our main package we have to do that with import <package_name>
*/

import "fmt"

// func keyword is used to define the functions

// syntax: func <function_name> (list of args to pass to the function) and inside the curly braces, we add the function definition.

func main() {
	fmt.Println("Hi There..!!")
}
