# TESTDLL

Objetivo deste projeto é testar a utilização do GO para criação de DLL na plataforma windows.

O projeto está dividido em três pacotes:
1. makeDLL: Cria DLL
2. includeDLL: Inclui DLL criada em 'makeDLL' no código executável
3. loadDLL: Carrega a  DLL criada em 'makeDLL'

[Referência](https://go.dev/wiki/WindowsDLLs).

## Requerimentos

- golang: https://go.dev/doc/install
- gcc: https://github.com/msys2/msys2-installer/

## Configurações

Go por padrão na plataforma windows não possui o CGO (C for GO) habilitado. Go compila C++ e é via esta funcionalidade que se cria DLL C++ para windows.  

Siga as próximas seções para configurar o ambiente para desenvolvimento GO + DLL Windows.

### Configurado GOPATH

Este passo não é obrigatório.   

Por padrão, trabalhe com GOPATH em 'C:\golang'.

1. Para ajustar, altere a variável ambiental do windows GOPATH=C:\golang
2. Execute o comando `go env -w GOPATH=C:\\golang`

Para confirmar utilize `go env GOPATH`.

### Habilitando CGO

A variável CGO_ENABLED por padrão é 0 na plataforma Windows.

Para habilitar, utilize o comando `go env -w CGO_ENABLED=1`

Para confirmar utilize `go env CGO_ENABLED`

[Documentação](https://pkg.go.dev/cmd/cgo).

## Instalando GCC

### Instruções GCC - minGW64 32bits

O projeto open source minGW34 mantém o suporte a 32bits com suporte a runtime *msvcrt* e *ucrt*.
Conforme [documentação](https://www.msys2.org/docs/environments/#msvcrt-vs-ucrt) o runtime *ucrt* está com melhor suporte e também é o default do Microsoft Visual Studio.

Página do projeto [mingGW64](https://github.com/niXman/mingw-builds-binaries?tab=readme-ov-file).

> Atualmente desenvolvendo suporte para aplicações 32bits

#### Build

Build para 32bits

`GOARCH=386 go build -o ./dist/makedll.dll -buildmode=c-shared main.go`

### Instruções GCC - MSYS2 64bits

É sugerida a instalação do GCC via [MSYS2](https://github.com/msys2/msys2-installer/). Baixe e instale a versão da sua plataforma.
Não é recomendável a instalação do ambiente C++ WinLibs ou MinGW por outras ferramentas - Podem conter trojans/vírus.
> As distribuições Linux já possuem as bibliotecas C++

O MSYS2 é comendado pela própria Microsoft. Infelizmente o Microsoft Visual C++ não suporta Go.

Neste [link oficial](https://code.visualstudio.com/docs/cpp/config-mingw) há o passo a passo da instalação e configuração do MSYS2, também reproduzido abaixo:

1. You can download the latest installer from the MSYS2 page or use this direct link to the installer.
2. Run the installer and follow the steps of the installation wizard. Note that MSYS2 requires 64 bit Windows 8.1 or newer.
3. In the wizard, choose your desired Installation Folder. Record this directory for later. In most cases, the recommended directory is acceptable. The same applies when you get to setting the start menu shortcuts step. When complete, ensure the Run MSYS2 now box is checked and select Finish. This will open a MSYS2 terminal window for you.
4. In this terminal, install the MinGW-w64 toolchain by running the following command:
`pacman -S --needed base-devel mingw-w64-ucrt-x86_64-toolchain`
5. Accept the default number of packages in the toolchain group by pressing Enter.
6. Enter Y when prompted whether to proceed with the installation.
7. Add the path of your MinGW-w64 bin folder to the Windows PATH environment variable by using the following steps:
   1. In the Windows search bar, type Settings to open your Windows Settings.
   2. Search for Edit environment variables for your account.
   3. In your User variables, select the Path variable and then select Edit.
   4. Select New and add the MinGW-w64 destination folder you recorded during the installation process to the list. If you used the default settings above, then this will be the path: C:\msys64\ucrt64\bin.
   5. Select OK, and then select OK again in the Environment Variables window to update the PATH environment variable. You have to reopen any console windows for the updated PATH environment variable to be available.

Check your MinGW installation: `gcc --version`.

## Pacotes

### makeDLL

Cria DLL com exibição de "HelloWorld".

Utilize o compila.bat para compilar e gerar a DLL.
O comando executado será `go build -o ./dist/makedll.dll -buildmode=c-shared main.go`.

### injectDLL

Este projeto *tenta* incluir a DLL no programa GO. Não obtive sucesso até o momento.

### loadDLL

Carrega a dll gerada em 'makeDLL' e executa a função HelloWorld.

Para executar execute o comando `go run main.go`.
