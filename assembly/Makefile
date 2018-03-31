.PHONY: run

run: clean 
	@/usr/local/bin/nasm -f macho64 hello_world.asm && ld -macosx_version_min 10.7.0 -lSystem -o hello_world hello_world.o && ./hello_world

clean: 
	@rm -rf hello_world.o hello_world
