## NASM

[NASM](https://en.wikipedia.org/wiki/Netwide_Assembler) - one of the most popular assembly languages.

Related resources for learning:

* NASM tutorials - http://cs.lmu.edu/~ray/notes/nasmtutorial/, https://www.youtube.com/watch?v=uca_zY8ZNpA
* LD - GNU linker, which creates executable file from the object file. More about linkers:  https://www.airs.com/blog/page/5?s=linker
* ELF - http://www.cirosantilli.com/elf-hello-world/

### Program layout 

```asm
; comment
section .data
	; stores constants 
	message: 	db "Hello World!", 10 ; assign label message to value Hello World!\n
	message_Len: 	equ $-message         ; store length of message

section .bss
	; stores uninitialized variables 

section .text
	; actual program
	global _start:

_start:
	; perform a system call	


```

### System calls

https://www.tutorialspoint.com/assembly_programming/assembly_system_calls.htm

or /usr/include/asm-generic



Example of printing to the screen and quit: 

```asm
...
_start:
	mov eax, 4
	mov ebx, 1
	mov ecx, msg,
	mov edx, msg_len
	systemcall

	; end the program
	mov eax, 1
	mov ebx, 0
	systemcall
```


### Cheatsheet

32-bit https://www.cs.uaf.edu/2008/fall/cs301/support/x86/index.html

Not-nasm assembly: https://cs.brown.edu/courses/cs033/docs/guides/x64_cheatsheet.pdf

https://www.cheatography.com/siniansung/cheat-sheets/linux-assembler/
