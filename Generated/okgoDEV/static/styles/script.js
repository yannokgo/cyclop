
// button light version
var down =  
document.getElementById('mainDIV'); 

function buttonFast(title) { 
    down.innerHTML = "<div class='hide' id='" + title + "' style='visibility:hidden;'><object data='./html/" + title + ".html' type='text/html'></object></div>"; 
    document.getElementById('mainDIV').innerHTML=document.getElementById(title).innerHTML;document.getElementById('checkbox').checked = false;
} 

function linkFast(title) { 
    down.innerHTML = "<div class='hide' id='" + title + "' style='visibility:hidden;'><object data='"+title+"' type='text/html'></object></div>"; 
    document.getElementById('mainDIV').innerHTML=document.getElementById(title).innerHTML;document.getElementById('checkbox').checked = false;
    } 
