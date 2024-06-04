/* Изменяющийся текст на главном экране */

function opacityMessage(){
  const firstMessage = document.getElementById("first-message")
  const secondMessage = document.getElementById("second-message")
  const thirdMessage = document.getElementById("third-message")
  secondMessage.style.display = "none";
  thirdMessage.style.display = "none";

  const messageOpacity = [
    {
      opacity: 0,
      easing: "ease-in",
    },
    {
      opacity: 1,
      easing: "ease-in"
    },
    {
      opacity: 0,
      easing: "ease-in",
    },
  ];

  const timing = {
    duration: 3000,
  };

  function relapseScript(){
    thirdMessage.style.display = "none";
    firstMessage.style.display = "block";
    firstMessage.animate(messageOpacity, timing);
    setTimeout(() => {
      firstMessage.style.display = "none";
      secondMessage.style.display = "block";
      secondMessage.animate(messageOpacity, timing);
    }, 3000)
    setTimeout(() => {
      secondMessage.style.display = "none";
      thirdMessage.style.display = "block";
      thirdMessage.animate(messageOpacity, timing);
    }, 6000)
  }

  relapseScript()

  setInterval(() => {
    relapseScript()
  }, 9000)

}

/* Popup Login */

function loginPopup(){
  const enter = document.querySelector(".circle");
  enter.addEventListener("click", function(){   // По клику на кнопку "Войти" окну добавляется класс "active", оно становится видимым. Фон затемняется.
    document.querySelector(".popup").classList.add("active");
    document.getElementById("overlay").style.display = "";
  });
  
  document.querySelector(".popup .close-btn").addEventListener("click", function(){ // По клику на кнопку "Крест" у окна удаляется класс "active", оно становится невидимым. Фон возвращается в стандартное состояние.
    document.querySelector(".popup").classList.remove("active");
    document.getElementById("overlay").style.display = "none";
  });
}

// Popup Logout 

function logOutPopup(){
  const circle = document.querySelector(".circle");
  circle.addEventListener("click", function(){  // По клику на кнопку "Вход" открывается всплывающее окно, фон размывается
    document.querySelector(".popup").classList.add("active");
    document.getElementById("overlay").style.display = "";
  });
  
  document.querySelector(".popup .close-btn").addEventListener("click", function(){ // По клику на кнопку "Крест" у окна удаляется класс "active", оно становится невидимым. Фон возвращается в стандартное состояние.
    document.querySelector(".popup").classList.remove("active");
    document.getElementById("overlay").style.display = "none";
  });

  circle.addEventListener('mouseover', function() {  // Если пользователь авторизован, то при наведении на иконку профиля появится крест
    document.getElementById("icon").setAttribute("class", "icon-cancel-1");
  });
    
  circle.addEventListener('mouseout', function() { 
    document.getElementById("icon").setAttribute("class", "icon-ok");
  });
}

function areYouSure(){
  const btn = document.querySelector("#plusOne");
  btn.addEventListener("click", function(){  // По клику на кнопку "Вход" открывается всплывающее окно, фон размывается
    document.querySelector(".popup-sure").classList.add("active");
    document.getElementById("overlay").style.display = "";
  });
  
  document.querySelector(".popup-sure .close-btn").addEventListener("click", function(){ // По клику на кнопку "Крест" у окна удаляется класс "active", оно становится невидимым. Фон возвращается в стандартное состояние.
    document.querySelector(".popup-sure").classList.remove("active");
    document.getElementById("overlay").style.display = "none";
  });
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
    } else if (score[i].innerHTML == 'н'){
      score[i].style.color = "brown"
    }
  }
}

/* Отслеживание выбранной роли на странице администратора для отображения ввода группы */

function changeRole(){
  var selectRole = document.getElementById("role");
  selectRole.addEventListener("click", function(){
    if (selectRole.value == "student"){
      document.getElementById("changeStudent").style.display = "none";
      document.getElementById("group").style.display = "inline-block";
    } else if (selectRole.value == "parent") {
      document.getElementById("changeStudent").style.display = "inline-block";
      document.getElementById("group").style.display = "none";
    } else if (selectRole.value == "teacher") {
      document.getElementById("changeStudent").style.display = "none";
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
    if (selectGroup.value != ""){  // Если выбирается селектор Группа, то селектор Преподаватель неактивен
      selectTeacher.setAttribute("disabled", true)
      selectTeacher.style.backgroundColor = "gray"
    } else if (selectTeacher.value != ""){ // Если выбирается селектор Преподаватель, то селектор Группа неактивен
      selectGroup.setAttribute("disabled", true)
      selectGroup.style.backgroundColor = "gray"
    }
  });
  selectForm.addEventListener('change', () => {  // Если селекторы сбрасываются на дефолт, то оба селектора активны
    if (selectGroup.value === "" && selectTeacher.value === ""){
      selectTeacher.removeAttribute("disabled", true)
      selectGroup.removeAttribute("disabled", true)
      selectGroup.style.backgroundColor = "white"
      selectTeacher.style.backgroundColor = "white"
    } else if (selectGroup.value === ""){
      selectTeacher.removeAttribute("disabled", true)
    } else if (selectTeacher.value === ""){
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

let nameStudent, date, scoreLevel, levelID, disciplineID;

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
              levelID = tab.rows[this.i].cells[this.j].attributes.level.value
              disciplineID = tab.rows[this.i].cells[this.j].attributes.discipline.value
              scoreScore = tab.rows[this.i].cells[this.j].innerHTML
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
              disciplineID = document.getElementById("dis").attributes.discipline.value
              scoreLevel = tab.rows[this.i].cells[this.j].innerHTML
              nameStudent = tab.rows[this.i].cells[0].innerHTML
              date = tab.rows[0].cells[this.j].innerHTML
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

/* Попап для журнала студента */

function checkLevelStnd(){
  var str = document.getElementsByClassName("score");
  for(let i = 0; i < str.length; i++){
    str[i].i = i;
    str[i].onclick=function(){
      let comment = str[this.i].attributes[3].value;
      document.getElementById("comment").innerHTML = comment;
      let score = str[this.i].innerHTML;
      document.getElementById("yourScore").innerHTML = score;
      let date = str[this.i].attributes[2].value;
      document.getElementById("date").innerHTML = date;
      popupForScore()
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

/* Добавление ДЗ и темы */

function checkScheduleTeacher(){
  let themes, homework;
  let tab = document.getElementById("table");
  for (let i = 1; i < tab.rows.length; i++){
    for (let j = 2; j < tab.rows[i].cells.length; j++){
        tab.rows[i].cells[j].i=i;
        tab.rows[i].cells[j].j=j;
        tab.rows[i].cells[j].onclick=function(){
          themes = tab.rows[this.i].cells[this.j].attributes[2].value;
          homework = tab.rows[this.i].cells[this.j].attributes[3].value;
          document.getElementById("ThemesHomeWork").setAttribute("lesID", tab.rows[this.i].cells[this.j].attributes[4].value);   
          document.getElementById("theme").innerHTML = themes;
          document.getElementById("homework").innerHTML = homework;
          popupForSchedule(tab.rows[this.i].cells[this.j].innerHTML)
        }
    }
  }
}

function checkScheduleStudent(){
  let themes, homework;
  let tab = document.getElementById("table");
  for (let i = 1; i < tab.rows.length; i++){
    for (let j = 2; j < tab.rows[i].cells.length; j++){
        tab.rows[i].cells[j].i=i;
        tab.rows[i].cells[j].j=j;
        tab.rows[i].cells[j].onclick=function(){
          themes = tab.rows[this.i].cells[this.j].attributes[2].value;
          homework = tab.rows[this.i].cells[this.j].attributes[3].value;
          document.getElementById("theme").innerHTML = themes;
          document.getElementById("homework").innerHTML = homework;
          popupForSchedule(tab.rows[this.i].cells[this.j].innerHTML)
        }
    }
  }
}

/* Слушатель для ячеек расписания */

function checkSchedule(){
  var tab=document.getElementById("table");
  for(let i = 1; i < tab.rows.length; i++){
    for(let j = 2; j < tab.rows[i].cells.length; j++){
      tab.rows[i].cells[j].i=i;
      tab.rows[i].cells[j].j=j;
      tab.rows[i].cells[j].onclick=function(){
        if (tab.rows[this.i].cells[this.j].innerHTML != ""){
          popupForSchedule(tab.rows[this.i].cells[this.j].innerHTML)
        }
      }
    }
  }
}

/* Цвет ячейки расписания */

function colorForSchedule(){
  let masLessons = document.getElementsByClassName("tableLessons")
  for (let i = 0; i < masLessons.length; i++){
    if (masLessons[i].attributes.typeLessons.value == "Лекция"){
      masLessons[i].style.backgroundColor = "#ffffcc"
    } else if (masLessons[i].attributes.typeLessons.value == "Семенар"){
      masLessons[i].style.backgroundColor = "#ccccff"
    }
  }
}

/* Подсчет Н-ок и средней оценки */ 

function avgSkips(){
  let countH = 0;
  let countLevel = 0;
  let summLevel = 0;
  let num = 0;
  let tab = document.getElementById("table")
  for(let i = 1; i < tab.rows.length; i++){
    for(let j = 1; j < tab.rows[i].cells.length-3; j++){
      if (tab.rows[i].cells[j].innerHTML == "н"){
        countH += 1;
      } else if(tab.rows[i].cells[j].innerHTML == "5" || tab.rows[i].cells[j].innerHTML == "4" || tab.rows[i].cells[j].innerHTML == "3" || tab.rows[i].cells[j].innerHTML == "2"){
        countLevel += 1
        summLevel += parseInt(tab.rows[i].cells[j].innerHTML)
      }
    }
    tab.rows[i].cells[tab.rows[i].cells.length-3].innerHTML = countH;
    if (countLevel == 0){
      tab.rows[i].cells[tab.rows[i].cells.length-2].innerHTML = ""
    } else{
      num = summLevel/countLevel;
      tab.rows[i].cells[tab.rows[i].cells.length-2].innerHTML = num.toFixed(2);
    }
    countH = 0;
    summLevel = 0;
    countLevel = 0;
  }
}

function avgSkipsStdt(){
  let countH = 0;
  let countLevel = 0;
  let summLevel = 0;
  let num = 0;
  let tab = document.getElementById("table")
  let stroke = document.getElementsByClassName("score");
  let n = 0;
  let a = ""
  let k = 0
  let count = 0;
  for(let i = 1; i < tab.rows.length; i++){
    if (tab.rows[i].cells[1].innerHTML != ""){
      a = tab.rows[i].cells[1].innerHTML
      while (a.indexOf("span", k) != -1){
        k = a.indexOf("span", k) + 1
        count += 1
      }
      console.log(count/2)
      for(let j = 0; j < count/2; j++){
        if (stroke[n].innerHTML == "н"){
          countH += 1;
          n += 1
        } else if(stroke[n].innerHTML == "5" || stroke[n].innerHTML == "4" || stroke[n].innerHTML == "3" || stroke[n].innerHTML == "2"){
          countLevel += 1
          summLevel += parseInt(stroke[n].innerHTML)
          n += 1
        }
      }
    }
    console.log(summLevel, countLevel, countH)
    tab.rows[i].cells[2].innerHTML = countH;
    if (countLevel == 0){
      tab.rows[i].cells[3].innerHTML = ""
    } else{
      num = summLevel/countLevel;
      tab.rows[i].cells[3].innerHTML = num.toFixed(2);
    }
    countH = 0;
    summLevel = 0;
    countLevel = 0;
    count = 0;
    a = ""
    k = 0
  }
}