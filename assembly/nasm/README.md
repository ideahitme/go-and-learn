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

How to perform system calls: https://www.tutorialspoint.com/assembly_programming/assembly_system_calls.htm

or /usr/include/asm-generic



Example of printing to the screen and quit: 

```asm
...
_start:
	mov eax, 4 ; see below for register naming
	mov ebx, 1
	mov ecx, msg,
	mov edx, msg_len
	systemcall ; for 64-bit CPU arch, for 32bit it is int 80h

	; end the program
	mov eax, 1
	mov ebx, 0
	systemcall
```


### Cheatsheet

32-bit https://www.cs.uaf.edu/2008/fall/cs301/support/x86/index.html

Not-nasm assembly: https://cs.brown.edu/courses/cs033/docs/guides/x64_cheatsheet.pdf

https://www.cheatography.com/siniansung/cheat-sheets/linux-assembler/

### Registers

The 16 integer registers are 64 bits wide and are called:
```
R0  R1  R2  R3  R4  R5  R6  R7  R8  R9  R10  R11  R12  R13  R14  R15
RAX RCX RDX RBX RSP RBP RSI RDI
```
(Note that 8 of the registers have alternate names.) You can treat the lowest 32-bits of each register as a register itself but using these names:

```
R0D R1D R2D R3D R4D R5D R6D R7D R8D R9D R10D R11D R12D R13D R14D R15D
EAX ECX EDX EBX ESP EBP ESI EDI
```
You can treat the lowest 16-bits of each register as a register itself but using these names:

```
R0W R1W R2W R3W R4W R5W R6W R7W R8W R9W R10W R11W R12W R13W R14W R15W
AX  CX  DX  BX  SP  BP  SI  DI
```
You can treat the lowest 8-bits of each register as a register itself but using these names:

```
R0B R1B R2B R3B R4B R5B R6B R7B R8B R9B R10B R11B R12B R13B R14B R15B
AL  CL  DL  BL  SPL BPL SIL DIL
```
For historical reasons, bits 15 through 8 of R0..R3 are named:

```
AH  CH  DH  BH
```

And finally, there are 16 XMM registers, each 128 bits wide, named:

XMM0 ... XMM15
Study this picture; hopefully it helps:


