<template>
 <div class="main-container">
     <div class="not-found d-flex justify-center align-center">
         <img src="@/assets/page_not_found.svg" alt="">
     </div>
 </div>
</template>

<style scoped>

html, body {
    overflow: hidden;
}
body {
  margin: 0;
  padding: 0;
}

.main-container {
  height: 100%;
  min-height: 100vh;
  width: 100vw;
  display: flex;
  justify-content: center;
  align-items: center;
  /* background: rgba(0, 0, 0, 0.1); */
  /* background: linear-gradient(40deg,
            rgba(252,104,45,1) 0%,
            rgba(39,5,85,1) 100%,
            rgba(242,5,87,1) 100%); */
  background: linear-gradient(135deg, #f8f8f8 0%,#efefef 100%);
}
img {
    width: 60%;
}

</style>
