<template>
  <header class="walpaper">
    <img src="@/assets/img/chartjs-logo.svg" />
    <h3>Budget Manager</h3>
  </header>
</template>

<script>
export default {
  name: "Header_",
};
</script>

<style scoped>
header {
  height: 30px;
  width: 100%;
  /* z-index: -1; */
  display: flex;
  font-family: 'quicksand bold';
  margin: 0px 0px;
  padding: 0px 0px;
  box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
}

header img {
  position: absolute;
  width: 30px;
  left: 5px;
}

h3 {
  text-align: center;
  width: 100%;
  margin: 7px 0px 0px 0px;
  color: #fa6976;
}
</style>
