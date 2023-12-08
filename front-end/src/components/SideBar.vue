<template>
  <div class="container sidebar">
    <div class="back">
      <img @click="clickBtnSideBar" src="@/assets/img/btn-slidebar.png" />
    </div>
    <div class="content">
      <Category />
      <AddExpense />
      <!-- <div class="back">
        <img @click="clickBtnSideBar" src="@/assets/img/btn-slidebar.png"/>
      </div> -->
    </div>
  </div>
</template>

<script>
import Category from "@/components/Category.vue";
import AddExpense from "./AddExpense/AddExpense.vue";

export default {
  name: "GraphView",
  components: {
    Category,
    AddExpense,
  },
  methods: {
    clickBtnSideBar() {
      document.querySelector(".sidebar").classList.toggle("close");
    },
  },
};
</script>

<style scoped>
.container.close {
  width: 30px;
}

.container {
  width: 300px;
  transition: 0.5s cubic-bezier(0.9, 0.59, 0.35, 1.15);;
  overflow: hidden;
  background-color: #ffff;
  font-family: "quicksand";
  font-weight: 600;
  height: 100%;
  color: black;
  box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
}

.container.close .content{
  opacity: 0;
}

.container .content{
  transition: .3s;
  height: 100%;
  opacity: 1;
}

div.back {
  width: 100%;
}
div.back:hover img {
  box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;

}
.container .back img {
  cursor: pointer;
  border-radius: 25px;
  transition: 0.7s;
  width: 25px;
}

.container.close .back img {
  transform: rotateZ(180deg);
}
</style>
