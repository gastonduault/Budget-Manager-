<template>
  <div class="AddExpense open">
    <div class="add-expense-revenue">
      <img @click="clickBack" src="@/assets/img/back.png" />
      <p>back</p>
    </div>
    <h3>Add Transaction</h3>
    <div class="form-add">
      <div class="form">
        <input
          v-on:keyup="enterLabel($event)"
          class="input-label"
          type="text"
          placeholder="Label"
        />
        <div class="input-number">
          <span @click="decrementPrice" class="input-number-decrement">–</span
          ><input
            v-on:keyup="enterPrice($event)"
            class="input-price"
            type="text"
            value="0"
          /><span @click="incrementPrice" class="input-number-increment"
            >+</span
          >
        </div>
        <div class="select-category">
          <button class="btn-category" v-if="subcategory == null">
            <img src="@/assets/img/category.png" />
            <p>Category</p>
          </button>
          <button class="btn-category" v-else>
            <img class="img-category" :src="maincategory" />
            <p>{{ subcategory }}</p>
          </button>
          <ListCategory @subcategory-selected="subcategorySelected" />
        </div>
        <div class="preview">
          <p class="preview-p">Preview</p>
          <div class="category-preview">
            <img v-if="maincategory !== null" :src="maincategory" />
            <div class="line">
              <div class="detail">
                <p v-if="subcategory !== null">{{ subcategory }}</p>
                <p class="label" v-if="label !== ''">({{ label }})</p>
              </div>
              <div class="price">{{ price }} 💲</div>
            </div>
          </div>
        </div>
        <button class="btn-add-transaction">ADD</button>
      </div>
    </div>
  </div>
</template>

<script>
import ListCategory from "@/components/AddExpense/ListCatgeory.vue";

export default {
  name: "AddExpense",
  components: {
    ListCategory,
  },
  data() {
    return {
      subcategory: null,
      maincategory: null,
      label: "",
      price: 0,
    };
  },
  methods: {
    // checkNumber(event) {
    //   let input = event.srcElement;
    //   input.style.animation = "none";
    //   if (
    //     !input.value.match(
    //       /^[-+]?[0-9]+\.[0-9]+$/
    //     )
    //   ) {
    //     input.value = "";
    //     input.style.animation = "shaking 3s";
    //   }
    // },
    enterLabel(event) {
      this.label = event.srcElement.value;
    },
    enterPrice(event) {
      this.price = event.srcElement.value;
    },
    subcategorySelected(data) {
      console.log("Esthétique et soins" == data.maincategory);
      switch (data.maincategory) {
        case "Logement":
          this.maincategory = require("@/assets/img/logement.png");
          break;
        case "Alimentation":
          this.maincategory = require("@/assets/img/alimentation.png");
          break;
        case "Auto et transport":
          console.log("passege");
          this.maincategory = require("@/assets/img/transport.png");
          break;
        case "Loisir et Sortie":
          this.maincategory = require("@/assets/img/loisir.png");
          break;
        case "Achat et Shopping":
          this.maincategory = require("@/assets/img/shopping.png");
          break;
        case "Banque":
          this.maincategory = require("@/assets/img/banque.png");
          break;
        case "Abonnement":
          this.maincategory = require("@/assets/img/abonnement.png");
          break;
        case "Divers":
          this.maincategory = require("@/assets/img/divers.png");
          break;
        case "Santé":
          this.maincategory = require("@/assets/img/sante.png");
          break;
        case "Dépense pro":
          this.maincategory = require("@/assets/img/depensePro.png");
          break;
        case "Esthétique et soins":
          this.maincategory = require("@/assets/img/esthetique.png");
          break;
        case "Impôt et taxe et soins":
          this.maincategory = require("@/assets/img/taxe.png");
          break;
        case "oper° Banque":
          this.maincategory = require("@/assets/img/cheque.png");
          break;
        case "Scolarité et enfant":
          this.maincategory = require("@/assets/img/enfant.png");
          break;
      }
      this.subcategory = data.subcategory;
    },
    clickBack() {
      document.querySelector(".category").classList.toggle("open");
      document.querySelector(".AddExpense").classList.toggle("open");
    },
    decrementPrice() {
      document.querySelector(".input-price").value--;
      this.price--;
    },
    incrementPrice() {
      this.price++;
      document.querySelector(".input-price").value++;
    },
    clickCategory() {
      document.querySelector(".select-category").classList.toggle(".open");
    },
  },
};
</script>

<style scoped>
.AddExpense {
  display: none;
}

.AddExpense.open {
  display: block;
}

.AddExpense {
  margin: 0px 0px;
  padding: 0px 0px;
  width: 100%;
  height: 100%;
  /* overflow: hidden; */
}

.AddExpense .add-expense-revenue {
  width: 100%;
  border-bottom: solid 1px black;
  display: flex;
  align-items: center;
  align-content: center;
  padding: 2px 0px;
}
.AddExpense .add-expense-revenue p {
  margin: 0px 0px;
}
.AddExpense .add-expense-revenue img {
  cursor: pointer;
  transition: 0.3s;
  width: 12px;
  height: 12px;
  margin-left: 5px;
  margin-right: 20px;
  border-radius: 25px;
  padding: 2px;
}
.AddExpense .add-expense-revenue img:hover {
  box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
  transform: translateX(5px);
}

div.form-add {
  margin: 5px 0px;
  /* display: flex;
  flex-direction: column; */
  height: 100%;
  overflow-y: auto;
}

.form-add .form {
  width: 100%;
  height: 88%;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
}

.form-add .form .input-label {
  display: block;
  background-color: #f2f2f2;
  padding: 10px;
  border: 0;
  margin: 0px auto 0px auto;
  font-size: 0.9em;
  font-family: "quicksand bold";
  border-radius: 25px 25px;
  width: 190px;
  box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
}

.form-add .form .input-number {
  display: block;
  width: 190px;
  margin: 0px auto;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  align-content: center;
}

@keyframes shaking {
  0% {
    transform: rotate(0deg);
  }
  25% {
    transform: rotate(5deg);
  }
  50% {
    transform: rotate(0eg);
  }
  75% {
    transform: rotate(-5deg);
  }
  100% {
    transform: rotate(0deg);
  }
}

.form-add .form .input-number input {
  box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
  font-size: 0.9em;
  font-family: "quicksand bold";
  padding: 8px;
  border: 0;
  border: solid 3px #f2f2f2;
  width: 130px;
}

.form-add .form .input-number span {
  width: 30px;
  font-family: "quicksand";
  background-color: #f2f2f2;
  text-align: center;
  padding: 10px;
  cursor: pointer;
}

.form-add .form .input-number span:hover {
  background-color: #dedede;
}

.form-add .form .input-number span:first-child {
  border-radius: 25px 0px 0px 25px;
}

.form-add .form .input-number span:last-child {
  border-radius: 0px 25px 25px 0px;
}

.btn-category {
  display: block;
  background-color: #e4e4e4;
  padding: 10px;
  border: 0;
  margin: 0px auto 0px auto;
  font-size: 0.9em;
  font-family: "quicksand bold";
  border-radius: 25px 25px 0px 0px;
  box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
  width: 236px;
  display: flex;
  align-items: center;
  cursor: pointer;
  transition: 0.3s;
}

.btn-category p {
  text-align: center;
  width: 100%;
  margin: 0px 25px 0px 0px;
}

.btn-category img {
  width: 20px;
  padding-left: 5px;
}

.preview {
  width: 100%;
  margin: 0px 0px;
}

.preview-p {
  text-align: center;
  margin: 0px 0px 0px 0px;
}

.category-preview {
  display: flex;
  list-style-type: none;
  justify-content: flex-start;
  align-content: center;
  align-items: center;
  width: 250px;
  margin: 0px auto;
  padding: 5px 0px;
  border: solid 1px #bbbbbb;
  border-radius: 5px 5px;
}

.category-preview img {
  width: 20px;
  height: 20px;
}

.category-preview .line p {
  margin: 0px 0px;
  padding: 0px 0px;
}

.category-preview .line {
  width: 100%;
  padding: 0px 0px 0px 0px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.category-preview .line .detail {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 0px 0px 0px 5px;
  margin: 0px 0px;
  font-size: 0.9em;
}
.category-preview .line .detail .label {
  margin-top: 5px;
}

.category-preview .price {
  color: #6ecd6e;
  font-family: "quicksand bold";
  font-size: 0.9em;
}

.btn-add-transaction {
  width: 190px;
  margin: 0px auto 15px auto;
  color: #fff;
  background-color: #5ad05a;
  font-family: "quicksand bold";
  font-size: 1em;
  padding: 8px 5px 5px 5px;
  border: none;
  border-radius: 25px 25px;
  box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
  transition: 0.3s;
  cursor: pointer;
}

.btn-add-transaction:hover {
  background-color: #55c655;
  transform: scale(1.1);
}
</style>
