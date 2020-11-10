
### xBindKeys
To install xbindkeys, and xbindkeys-config (the GUI for xbindkeys)

```bash
sudo apt-get install xbindkeys xbindkeys-config 
```

To create the default config file for xbindkeys

```bash
xbindkeys --defaults > $HOME/.xbindkeysrc  
```

Then run:
```bash
xbindkeys
xbindkeys-config
```
And for your first keybinding, you may find it useful to assign Ctrl+Shift+Alt+X, or whatever you prefer, to xbindkeys-config

To keep the xbindkeys hotkeys active ever time you start the computer...
```
Main Menu
  System  
    Preferences  
      Startup Applications  
        [ Add ]  
          Name:     xbindkeys  
          Command:  xbindkeys  
          Comment:  xbindkeys  
```

## HTML to JS
```html
<script>

// HTML to JavaScript converter
// By John Krutsch (http://asp.xcentrixlc.com/john)
// Moderator of the JavaScript Kit Help Forum: http://freewarejava.com/cgi-bin/Ultimate.cgi

function scriptIt(val){
val.value=val.value.replace(/"/gi,"&#34;")
val.value=val.value.replace(/'/gi,"&#39;")
valArr=escape(val.value).split("%0D%0A")
val.value=""
for (i=0; i<valArr.length; i++){
val.value+= (i==0) ? "<script>\ninfo=" : ""
val.value+= "\"" + unescape(valArr[i])
val.value+= (i!=valArr.length-1) ? "\" + \n" : "\"\n" 
}
val.value+="\ndocument.write(info)\n<\/script>"
}

function ctrlA(el) {
with(el){
focus(); select() 
}
if(document.all){
txt=el.createTextRange()
txt.execCommand("Copy") 
window.status='Selected and copied to clipboard!'
}
else window.status='Press ctrl-c to copy the text to the clipboard'
setTimeout("window.status=''",3000)
} 

</script>
<center>
<form name="f">
<input type="button" value="HTML -> JavaScript" onclick="scriptIt(document.f.t)" />
<input type="reset" value="Reset">
<input type="button" value="Select All" onclick="ctrlA(document.f.t)" />
<br />
<textarea name="t" cols="75"
rows="20"></textarea>
<br />
</form>
</center>
```