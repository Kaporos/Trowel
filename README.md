# Trowel

Trowel is a **really** lightweight project scaffolding tool.

# Usage

First, you must create your templates, so open a dir and put some files in it.

    mkdir great-template
    echo "print('hello world')" > great-template/main.py

So here we have the most basic python app.

Then, register the template using trowel

    trowel register great-template

This command will make a copy of it inside of `$HOME/.local/share/trowel`
(which means after that you can safely delete the great-template if you want to).

After registering some template, you can create a new project using

    trowel new great-template my-new-project


And there you go ! `my-new-project` was just created !

If you don't specify any name, then the project will be created in `/tmp` directory.
    
