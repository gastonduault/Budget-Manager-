<template>
  <main>
    <div class="main">
      <Header_/>
      <Graph />
    </div>
    <SideBar />
  </main>
</template>
<script>
import Header_ from "@/components/Header_.vue";
import Graph from "@/components/Graph.vue";
import SideBar from "@/components/SideBar.vue";

export default {
  name: "DisplayBudget",
  components: {
    Graph,
    SideBar,
    Header_
  },
};
</script>

<style scoped>
main{
  display: flex;
  width: 100vw;
  height: 100vh;
  margin: 0px;
  padding: 0px;
  overflow: hidden;
}

.main {
  margin: 0px 0px;
  padding: 0px 0px;
  height: 100%;
  width: 100%;
}
</style>
