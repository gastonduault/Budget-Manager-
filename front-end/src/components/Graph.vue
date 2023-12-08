<template>
  <div class="container">
    <div class="limitChart">
      <Bar id="my-chart-id" :options="chartOptions" :data="chartData" />
    </div>
  </div>
</template>

<script>
import { Bar } from "vue-chartjs";
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  BarElement,
  CategoryScale,
  LinearScale,
  Colors
} from "chart.js";

ChartJS.register(
  Title,
  Tooltip,
  BarElement,
  CategoryScale,
  LinearScale,
  Colors
);

export default {
  name: "GraphView",
  components: { Bar },
  data() {
    return {
      chartData: {
        labels: [
          "Logement",
          "Alimentaire",
          "Auto & Transport",
          "Banque",
          "Abonnement",
          "Divers",
          "Santé",
          "Scolarité enfant"
        ],
        datasets: [
          {
            data: [500, 160, 111, 7, 22, 47, 34, 48.75],
            Colors
          },
        ],
      },
      chartOptions: {
        responsive: true,
      },
    };
  },
};
</script>

<style scoped>
.container {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
}
.limitChart{
  width: 80%;
  height: auto;
  margin: 0px auto;
}
</style>
