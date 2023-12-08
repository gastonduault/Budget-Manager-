<template>
  <router-view/>
</template>

<style>
@font-face {
    font-family: 'quicksand';
    src: url('@/assets/fonts/quicksand/Quicksand_Light.otf');
    font-weight: normal;
    font-style: normal;
}

@font-face {
    font-family: 'quicksand bold';
    src: url('@/assets/fonts/quicksand/Quicksand_Bold.otf');
    font-weight: normal;
    font-style: normal;
}

body {
  margin: 0px 0px;
  padding: 0px 0px;
}

::-webkit-scrollbar {
  background-color: transparent;
  width: 5px;
}

::-webkit-scrollbar-track {
  background-color: transparent;
}

::-webkit-scrollbar-thumb {
  background-color: #888888;
  border-radius: 5px 5px;
  width: 0.5em !important;
}

</style>
