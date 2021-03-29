<template>
      <router-view/>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Signika:wght@300;400&display=swap');

html, body {
    overflow: hidden;
    color: #fff;
}
body {
  margin: 0;
  padding: 0;
}
hr {
  margin: 10px 0px !important;
  border-top: 2px solid #FFFFFF33 !important;
}

h2, h1 {
  color: #ccc; 
}

#app {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;

  font-family: 'Signika', sans-serif;
  /* font-family: Avenir, Helvetica, Arial, sans-serif; */
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* background: linear-gradient(0deg, #50e3c2 0%,#10d574 100%); */
  /* background: rgba(0,0,0,0.01); */
  /* background: #50e3c2; */
  /* background: #121212; */
  background: #1E1E1E;
  /* background: linear-gradient(135deg, #50e3c2 0%,#10d574 100%); */
  /* background-image: url("assets/dash_bottom.svg"), url("assets/dash_top.svg");
  background-repeat: no-repeat, no-repeat;
  background-position: left bottom, right top; */
  /* background-image: url("assets/dash_bottom_green.svg"), url("assets/dash_top_green.svg"); */
  background-image: url("assets/dash_bottom_green.svg"), url("assets/dash_top_green.svg");
  background-repeat: no-repeat, no-repeat;
  background-position: left bottom, right top;
  background-size: contain, contain;
}
a {
  text-decoration: underline;
  background: linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.scroll_view {
  overflow-y: scroll;
  height: 800px;
  min-height: 400px;
  max-height: 800px;
}
.text-capture {
  background: linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.text-capture.big {
  font-size: 20px;
}
.btn-standard {
  background: transparent;
  border: 1px solid #50e3c2;
  color: #10d574;
  border-radius: 8px;
  padding: 2px 5px;
  outline: none;
}

.btn-standard:focus {
  outline: none;
}
.btn-standard:hover {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background: linear-gradient(0deg, #50e3c2 0%,#10d574 100%);
  color: #121212;
}

.info-text {
  color: #fff;
}

.tag-list {
  display: flex;
  flex-wrap: wrap;
}
.tag-standard {
  padding: 3px;
  margin: 0 3px;
  background: linear-gradient(0deg, #50e3c2 0%,#10d574 100%);
  color: #121212;
  border-radius: 4px;
  width: min-content;
  font-weight: bolder;
}

.danger-text {
  color: rgba(255, 0, 0, 1);
}
.danger-text small {
  color: rgba(255, 255, 255, 0.8);
}

.form-control:disabled {
  background-color: #30363D !important;
}
.main-view {
    display: grid;
    grid-template-columns: auto;
    grid-template-rows: 50px 1fr;
    grid-column-gap: 10px;
    grid-row-gap: 10px;
    
    min-width: 650px;
    height: 90%;
    /* background: rgb(255,255,255,0.8); */
    padding: 10px;
}
.box_component {
   border-radius: 8px;
   box-shadow: 0px 0px 5px 2px rgb(0,0,0, 0.1);
   padding: 15px;
   background: #fff;
}

.circle-img {
  border-radius: 50%;
  border: 2px solid #30363D;
}
.circle-img.medium {
  object-fit: cover;
  width: 55px;
  height: 55px;
}
.circle-img.big {
  object-fit: cover;
  width: 100px;
  height: 100px;
}

input {
  color: #30363D;
}
/* 
  UTIL CLASSES
 */
.d-grid {
  display: grid;
}
.d-flex {
  display: flex;
}
.justify-between {
  justify-content: space-between;
}
.justify-center {
  justify-content: center;
}
.justify-evenly {
  justify-content: space-evenly;
}
.justify-start {
  justify-content: flex-start;
}
.justify-end {
  justify-content: flex-end;
}
.align-center {
  align-items: center;
}
.flex-wrap {
  flex-wrap: wrap;
}
.icon.green {
  background:linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.standard-font {
  font-size: 14px;
}
.icon.white {
  color: #fff;
}
.icon.big {
  font-size: 20px;
}
.icon.medium {
  font-size: 16px;
}

.icon.hover:hover {
  cursor: pointer;
}
.btn-normal {
  background: #fff;
  background:linear-gradient(135deg, #50e3c2 0%,#10d574 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  transition: all 5s ease-in-out;
  border-radius: 8px;
}
.btn-normal:focus,
.btn-normal:hover {
  -webkit-text-fill-color: #fff;
  background: linear-gradient(135deg, #50e3c2 0%,#10d574 100%);;
}

.icon-delete:hover, .icon-user-minus:hover {
  -webkit-text-fill-color: rgba(255, 0, 0,1);
}

.btn.no-line-height {
  line-height: 0 !important;
}

.dots {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* MODAL STUFF */
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  border-radius: 8px;
  width: 450px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  font-family: Helvetica, Arial, sans-serif;
}

.modal-header h3 {
  margin-top: 0;
  color: #42b983;
}

.modal-body {
  margin: 20px 0;
}

.modal-default-button {
  float: right;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}

hr {
  margin: 5px 0px;
}
</style>
