
set nu
set cursorline
set hlsearch
set ts=4
set autoindent
set ignorecase

let g:airline#extensions#tabline#enabled = 1

call plug#begin('/home/siuyut/.vim/plugged')
Plug 'mhinz/vim-startify'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'Yggdroot/indentLine'

Plug 'w0ng/vim-hybrid'
Plug 'rakr/vim-one'
Plug 'KeitaNakamura/neodark.vim'

Plug 'Raimondi/delimitMate'
call plug#end()

set termguicolors
colorscheme neodark
set background=dark
