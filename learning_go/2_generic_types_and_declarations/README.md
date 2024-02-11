# 2_generic_types_and_declarations

* 5 types of literals in Go: 
    * integer
    * floating-point
    * rune
    * string
    * complex

* differences between `var` and `:=`:
    * `var` - declares a variable, but does not initialize it
    * `:=` - declares and initializes a variable
* more subtle: `:=` should (and can) only be used inside function body, while `var` also works on package-level (though, it's rare to declare variable like this - declarations list should then be used)
* sometimes `:=` shouldn't be used inside function body (e.g. when you want to declare a variable with the same name as a package-level variable) to make your intention clear
* `const` - declares a constant, which must be initialized at the time of declaration
* `namingConvention`: 
    * `PascalCase` - for exported names (visible outside of package)
    * `camelCase` - for internal names (visible only inside package)
    * no `snake_case` in Go as in Python
