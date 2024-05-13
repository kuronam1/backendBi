/* Изменяющийся текст на главном экране */

$(document).ready(function(){
        $(".welcome-message .rotate").textrotator({
    animation: "fade",
    speed: 1000
    });
});

!function($){
  
    var defaults = {
          animation: "dissolve",
          separator: ",",
          speed: 2000
      };
      
      $.fx.step.textShadowBlur = function(fx) {
      $(fx.elem).prop('textShadowBlur', fx.now).css({textShadow: '0 0 ' + Math.floor(fx.now) + 'px black'});
    };
      
    $.fn.textrotator = function(options){
      var settings = $.extend({}, defaults, options);
      
      return this.each(function(){
        var el = $(this)
        var array = [];
        $.each(el.text().split(settings.separator), function(key, value) { 
          array.push(value); 
        });
        el.text(array[0]);
        
        // animation option
        var rotate = function() {
          switch (settings.animation) { 
            case 'dissolve':
              el.animate({
                textShadowBlur:20,
                opacity: 0
              }, 500 , function() {
                index = $.inArray(el.text(), array)
                if((index + 1) == array.length) index = -1
                el.text(array[index + 1]).animate({
                  textShadowBlur:0,
                  opacity: 1
                }, 500 );
              });
            break;
            
            case 'flip':
              if(el.find(".back").length > 0) {
                el.html(el.find(".back").html())
              }
            
              var initial = el.text()
              var index = $.inArray(initial, array)
              if((index + 1) == array.length) index = -1
              
              el.html("");
              $("<span class='front'>" + initial + "</span>").appendTo(el);
              $("<span class='back'>" + array[index + 1] + "</span>").appendTo(el);
              el.wrapInner("<span class='rotating' />").find(".rotating").hide().addClass("flip").show().css({
                "-webkit-transform": " rotateY(-180deg)",
                "-moz-transform": " rotateY(-180deg)",
                "-o-transform": " rotateY(-180deg)",
                "transform": " rotateY(-180deg)"
              })
              
            break;
            
            case 'flipUp':
              if(el.find(".back").length > 0) {
                el.html(el.find(".back").html())
              }
            
              var initial = el.text()
              var index = $.inArray(initial, array)
              if((index + 1) == array.length) index = -1
              
              el.html("");
              $("<span class='front'>" + initial + "</span>").appendTo(el);
              $("<span class='back'>" + array[index + 1] + "</span>").appendTo(el);
              el.wrapInner("<span class='rotating' />").find(".rotating").hide().addClass("flip up").show().css({
                "-webkit-transform": " rotateX(-180deg)",
                "-moz-transform": " rotateX(-180deg)",
                "-o-transform": " rotateX(-180deg)",
                "transform": " rotateX(-180deg)"
              })
              
            break;
            
            case 'flipCube':
              if(el.find(".back").length > 0) {
                el.html(el.find(".back").html())
              }
            
              var initial = el.text()
              var index = $.inArray(initial, array)
              if((index + 1) == array.length) index = -1
              
              el.html("");
              $("<span class='front'>" + initial + "</span>").appendTo(el);
              $("<span class='back'>" + array[index + 1] + "</span>").appendTo(el);
              el.wrapInner("<span class='rotating' />").find(".rotating").hide().addClass("flip cube").show().css({
                "-webkit-transform": " rotateY(180deg)",
                "-moz-transform": " rotateY(180deg)",
                "-o-transform": " rotateY(180deg)",
                "transform": " rotateY(180deg)"
              })
              
            break;
            
            case 'flipCubeUp':
              if(el.find(".back").length > 0) {
                el.html(el.find(".back").html())
              }
            
              var initial = el.text()
              var index = $.inArray(initial, array)
              if((index + 1) == array.length) index = -1
              
              el.html("");
              $("<span class='front'>" + initial + "</span>").appendTo(el);
              $("<span class='back'>" + array[index + 1] + "</span>").appendTo(el);
              el.wrapInner("<span class='rotating' />").find(".rotating").hide().addClass("flip cube up").show().css({
                "-webkit-transform": " rotateX(180deg)",
                "-moz-transform": " rotateX(180deg)",
                "-o-transform": " rotateX(180deg)",
                "transform": " rotateX(180deg)"
              })
              
            break;
            
            case 'spin':
              if(el.find(".rotating").length > 0) {
                el.html(el.find(".rotating").html())
              }
              index = $.inArray(el.text(), array)
              if((index + 1) == array.length) index = -1
              
              el.wrapInner("<span class='rotating spin' />").find(".rotating").hide().text(array[index + 1]).show().css({
                "-webkit-transform": " rotate(0) scale(1)",
                "-moz-transform": "rotate(0) scale(1)",
                "-o-transform": "rotate(0) scale(1)",
                "transform": "rotate(0) scale(1)"
              })
            break;
            
            case 'fade':
              el.fadeOut(settings.speed, function() {
                index = $.inArray(el.text(), array)
                if((index + 1) == array.length) index = -1
                el.text(array[index + 1]).fadeIn(settings.speed);
              });
            break;
          }
        };
        setInterval(rotate, settings.speed);
      });
    }
    
  }(window.jQuery);



/* Popup Login */

function loginPopup(){
  const enter = document.querySelector("#entry");
  enter.addEventListener("click", function(){   // По клику на кнопку "Войти" окну добавляется класс "active", оно становится видимым. Фон затемняется.
    document.querySelector(".popup").classList.add("active");
    document.getElementById("overlay").style.display = "";
  });
  
  document.querySelector(".popup .close-btn").addEventListener("click", function(){ // По клику на кнопку "Крест" у окна удаляется класс "active", оно становится невидимым. Фон возвращается в стандартное состояние.
    document.querySelector(".popup").classList.remove("active");
    document.getElementById("overlay").style.display = "none";
  });
}

// Popup Log out 

function logOutPopup(){
  const authorized = document.getElementById("authorized");
  authorized.addEventListener("click", function(){  // По клику на кнопку "Вход" открывается всплывающее окно, фон размывается
    document.querySelector(".popup").classList.add("active");
    document.getElementById("overlay").style.display = "";
  });
  
  document.querySelector(".popup .close-btn").addEventListener("click", function(){ // По клику на кнопку "Крест" у окна удаляется класс "active", оно становится невидимым. Фон возвращается в стандартное состояние.
    document.querySelector(".popup").classList.remove("active");
    document.getElementById("overlay").style.display = "none";
  });
  
  const circleExit = document.getElementById("circle");

  circleExit.addEventListener('mouseover', function() {  // Если пользователь авторизован, то при наведении на иконку профиля появится крест
    document.getElementById("icon").setAttribute("class", "icon-cancel-1");
  });
    
  circleExit.addEventListener('mouseout', function() { 
    document.getElementById("icon").setAttribute("class", "icon-ok");
  });
}

/* Изменяющиеся стили для кнопок в хедере  */

function restyleLinks(){
  let substringSchedule = "schedule";
  let substringJournal = "journal";

  // Все работает от входящих в URL слов. Если изменятся url'ы, то необходимо поменять условия ниже

  if (document.URL.includes(substringSchedule)){ // Для расписания
    let link = document.getElementById("schedule");
    link.setAttribute("onclick", "return false");
    link.style.boxShadow = "0px 0px 5px var(--color1)";
    link.style.cursor = "default";
  } else if (document.URL.includes(substringJournal)){  // Для журнала
    let link = document.getElementById("journal");
    link.setAttribute("onclick", "return false");
    link.style.boxShadow = "0px 0px 5px var(--color1)";
    link.style.cursor = "default";
  }
}

function restyleLinksStudent(){
  let substringStudentSchedule = "schedule";
  let suubstringStudentJournal = "journal";

  // Все работает от входящих в URL слов. Если изменятся url'ы, то необходимо поменять условия ниже

  if (document.URL.includes(substringStudentSchedule)){ // Для расписания
    let link = document.getElementById("student-schedule");
    link.setAttribute("onclick", "return false");
    link.style.boxShadow = "0px 0px 5px var(--color1)";
    link.style.cursor = "default";
  } else if (document.URL.includes(suubstringStudentJournal)){  // Для журнала
    let link = document.getElementById("student-journal");
    link.setAttribute("onclick", "return false");
    link.style.boxShadow = "0px 0px 5px var(--color1)";
    link.style.cursor = "default";
  }
}

function restyleLinksTeacher(){
  let substringTeacherSchedule = "schedule";
  let suubstringTeacherJournal = "journal";

  // Все работает от входящих в URL слов. Если изменятся url'ы, то необходимо поменять условия ниже

  if (document.URL.includes(substringTeacherSchedule)){ // Для расписания
    let link = document.getElementById("teacher-schedule");
    link.setAttribute("onclick", "return false");
    link.style.boxShadow = "0px 0px 5px var(--color1)";
    link.style.cursor = "default";
  } else if (document.URL.includes(suubstringTeacherJournal)){  // Для журнала
    let link = document.getElementById("teacher-journal");
    link.setAttribute("onclick", "return false");
    link.style.boxShadow = "0px 0px 5px var(--color1)";
    link.style.cursor = "default";
  }
}

/* Цвет для оценок в журнале админа и учителя */

function changeColor(){
  let score = document.getElementsByClassName("score");
  for (let i = 0; i < score.length; i++){
    if (score[i].innerHTML == 5){
      score[i].style.color = "green"
    } else if (score[i].innerHTML == 4){
      score[i].style.color = "#cccc00"
    } else if (score[i].innerHTML == 3){
      score[i].style.color = "orange"
    } else if (score[i].innerHTML == 2){
      score[i].style.color = "red"
    } else if (score[i].innerHTML == 'Н'){
      score[i].style.color = "brown"
    }
  }
}

/* Отслеживание выбранной роли на странице администратора для отображения ввода группы */

function changeRole(){
  var selectRole = document.getElementById("role");
  selectRole.addEventListener("click", function(){
    if (selectRole.value == "student" || selectRole.value == "parent"){
      document.getElementById("group").style.display = "inline-block";
    } else {
      document.getElementById("group").style.display = "none";
    }
  })
}

/* Выбор для журнала Администратора на один селектор */ 

function oneSelector(){
  let selectGroup = document.getElementById("group")
  let selectTeacher = document.getElementById("teacher")
  let selectForm = document.getElementById("selectForm")
  
  selectForm.addEventListener('change', () => {
    if (selectGroup.value != "false"){  // Если выбирается селектор Группа, то селектор Преподаватель неактивен
      selectTeacher.setAttribute("disabled", true)
      selectTeacher.style.backgroundColor = "gray"
    } else if (selectTeacher.value != "0"){ // Если выбирается селектор Преподаватель, то селектор Группа неактивен
      selectGroup.setAttribute("disabled", true)
      selectGroup.style.backgroundColor = "gray"
    }
  });

  selectForm.addEventListener('change', () => {  // Если селекторы сбрасываются на дефолт, то оба селектора активны
    if (selectGroup.value === "false" && selectTeacher.value === "false"){
      selectTeacher.removeAttribute("disabled", true)
      selectGroup.removeAttribute("disabled", true)
      selectGroup.style.backgroundColor = "white"
      selectTeacher.style.backgroundColor = "white"
    } else if (selectGroup.value === "0"){
      selectTeacher.removeAttribute("disabled", true)
    } else if (selectTeacher.value === "0"){
      selectGroup.removeAttribute("disabled", true)
    }
  });
  
}

/* Окно для оценок */

function popupForScore(){
  document.querySelector(".popupScore").classList.add("active"); 
  document.getElementById("overlay").style.display = "";
  
  document.querySelector(".popupScore .close-btn").addEventListener("click", function(){ // По клику на кнопку "Крест" у окна удаляется класс "active", оно становится невидимым. Фон возвращается в стандартное состояние.
    document.querySelector(".popupScore").classList.remove("active");
    document.getElementById("overlay").style.display = "none";
  });

}


/* Выставление оценок в всплывающем окне */

let nameStudent, date;

function setScoreInPopup(role){
  var tab=document.getElementById("table");
  for(let i = 1; i < tab.rows.length; i++){
    for(let j = 1; j < tab.rows[i].cells.length; j++){
      if (j != tab.rows[i].cells.length - 2 && j != tab.rows[i].cells.length - 3){
        tab.rows[i].cells[j].i=i;
        tab.rows[i].cells[j].j=j;
        tab.rows[i].cells[j].onclick=function(){
          if (role === "adm"){ // Тут должна быть проверка роли - учитель или админ
            if (tab.rows[this.i].cells[this.j].innerHTML != ''){ // Вариант, когда оценка стоит, т.е. её можно только изменить (для админа)
              nameStudent = '';
              date = "";
              nameStudent = tab.rows[this.i].cells[0].innerHTML
              date = tab.rows[0].cells[this.j].innerHTML
              popupForScore()
            } else {
              console.log("Нельзя добавить новую оценку!")
            }
          } else if (role === "tchr") { // Для учителя
            if (tab.rows[this.i].cells[this.j].innerHTML === ''){ // Вариант, когда оценка не стоит, т.е. её можно только проставить (для учителя)
              nameStudent = '';
              date = "";
              nameStudent = tab.rows[this.i].cells[0].innerHTML
              date = tab.rows[0].cells[this.j].innerHTM
              popupForScore()
            } else {
              console.log("Нельзя изменить оценку!")
            }
          }
        }
      }
    }
  }
}

/* Окно для расписания */

function popupForSchedule(data){
  document.querySelector(".popupSchedule").classList.add("active");
  document.getElementById("overlay").style.display = "";
  document.getElementById("windowSchedule").innerHTML = data;

  document.querySelector(".popupSchedule .close-btn").addEventListener("click", function(){ // По клику на кнопку "Крест" у окна удаляется класс "active", оно становится невидимым. Фон возвращается в стандартное состояние.
    document.querySelector(".popupSchedule").classList.remove("active");
    document.getElementById("overlay").style.display = "none";
  });
}

/* Слушатель для ячеек расписания */

function checkSchedule(){
  var tab=document.getElementById("table");
  for(let i = 1; i < tab.rows.length; i++){
    for(let j = 2; j < tab.rows[i].cells.length; j++){
      tab.rows[i].cells[j].i=i;
      tab.rows[i].cells[j].j=j;
      tab.rows[i].cells[j].onclick=function(){
        popupForSchedule(tab.rows[this.i].cells[this.j].innerHTML)
      }
    }
  }
}
