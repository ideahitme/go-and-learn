## Linker

Source: https://www.airs.com/blog/archives/39

Linker operates on three data types (found in object files): 

1. Symbol - name and a value. There are two types of symbols: 

- Defined symbol - defines symbols existing in the same object file, e.g. each function, global and static variable. Value of each symbol is an offset in the content (see below). Linker then assign a memory address to each symbol. 	

- Undefined symbol - is a reference to a value in another object file. Linker will resolve the address of such symbols. 

2. Relocation - is a computation to be performed on the given offset in the content. Normally relocation includes a symbol, and optionally and operand known as `addend`. For example, relocation may assigned a value for symbol `A`, which is equal to symbol `B` plus some constant value.  

3. Content - this is what memory should look like during execution time. They contain the values of initialized data, unitialized data, constants etc. Linker reads this data and applies relocation to it and writes the result into an executable file. 


Linker Steps 
1. Read the input object files. Determine the length and type of the contents. Read the symbols.
2. Build a symbol table containing all the symbols, linking undefined symbols to their definitions.
3. Decide where all the contents should go in the output executable file, which means deciding where they should go in memory when the program runs.
4. Read the contents data and the relocations. Apply the relocations to the contents. Write the result to the output file.
5. Optionally write out the complete symbol table with the final values of the symbols.


