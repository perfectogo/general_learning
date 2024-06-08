
; You may customize this and other start-up templates; 
; The location of this template is c:\emu8086\inc\0_com_template.txt

ORG 100h

    PUSH 0FFFFh
    PUSH 32h
    
    POP AX
    POP CX

RET
   
   
END
 