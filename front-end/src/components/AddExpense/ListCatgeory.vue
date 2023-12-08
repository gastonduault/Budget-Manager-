<template>
  <div class="container">
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/logement.png" />
        <p class="p-category-title">Logement</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">
          Assurance Habitation
        </p>
        <p @click="clickSubcategory($event)" class="subcategory">Loyer</p>
        <p @click="clickSubcategory($event)" class="subcategory">Charges</p>
        <p @click="clickSubcategory($event)" class="subcategory">Eau</p>
        <p @click="clickSubcategory($event)" class="subcategory">Electricité</p>
        <p @click="clickSubcategory($event)" class="subcategory">Entretien</p>
        <p @click="clickSubcategory($event)" class="subcategory">Gaz</p>
        <p @click="clickSubcategory($event)" class="subcategory">Logement - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/alimentation.png" />
        <p class="p-category-title">Alimentation</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Restaurant</p>
        <p @click="clickSubcategory($event)" class="subcategory">Fast foods</p>
        <p @click="clickSubcategory($event)" class="subcategory">Supermarché</p>
        <p @click="clickSubcategory($event)" class="subcategory">Café</p>
        <p @click="clickSubcategory($event)" class="subcategory">Almientat - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/transport.png" />
        <p class="p-category-title">Auto et transport</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Entretien véhicule</p>
        <p @click="clickSubcategory($event)" class="subcategory">Transport commun</p>
        <p @click="clickSubcategory($event)" class="subcategory">Carburant</p>
        <p @click="clickSubcategory($event)" class="subcategory">Location</p>
        <p @click="clickSubcategory($event)" class="subcategory">Assurance</p>
        <p @click="clickSubcategory($event)" class="subcategory">Billet avion</p>
        <p @click="clickSubcategory($event)" class="subcategory">Billet train</p>
        <p @click="clickSubcategory($event)" class="subcategory">Péage</p>
        <p @click="clickSubcategory($event)" class="subcategory">Stationnement</p>
        <p @click="clickSubcategory($event)" class="subcategory">Transport - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/loisir.png" />
        <p class="p-category-title">Loisir <em></em> Sortie</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Bar / Club</p>
        <p @click="clickSubcategory($event)" class="subcategory">Divertissement</p>
        <p @click="clickSubcategory($event)" class="subcategory">Frais Animaux</p>
        <p @click="clickSubcategory($event)" class="subcategory">Hobbies</p>
        <p @click="clickSubcategory($event)" class="subcategory">Hôtels</p>
        <p @click="clickSubcategory($event)" class="subcategory">Sport</p>
        <p @click="clickSubcategory($event)" class="subcategory">Voyage / Vancaces</p>
        <p @click="clickSubcategory($event)" class="subcategory">Loisir - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/shopping.png" />
        <p class="p-category-title">Achat <em></em> Shopping</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Cadeaux</p>
        <p @click="clickSubcategory($event)" class="subcategory">Culturel</p>
        <p @click="clickSubcategory($event)" class="subcategory">
          Vêtement / Chaussure
        </p>
        <p @click="clickSubcategory($event)" class="subcategory">Hight Tech</p>
        <p @click="clickSubcategory($event)" class="subcategory">Licence</p>
        <p @click="clickSubcategory($event)" class="subcategory">Achat - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/banque.png" />
        <p class="p-category-title">Banque</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Incident de paiment</p>
        <p @click="clickSubcategory($event)" class="subcategory">Frais bancaires</p>
        <p @click="clickSubcategory($event)" class="subcategory">Débit mensuel</p>
        <p @click="clickSubcategory($event)" class="subcategory">Epargne</p>
        <p @click="clickSubcategory($event)" class="subcategory">Hypothèque</p>
        <p @click="clickSubcategory($event)" class="subcategory">
          Remboursement emprunt
        </p>
        <p @click="clickSubcategory($event)" class="subcategory">Service Bancaires</p>
        <p @click="clickSubcategory($event)" class="subcategory">Banque - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/abonnement.png" />
        <p class="p-category-title">Abonnement</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Internet</p>
        <p @click="clickSubcategory($event)" class="subcategory">Téléphone</p>
        <p @click="clickSubcategory($event)" class="subcategory">Abonnement - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/divers.png" />
        <p class="p-category-title">Divers</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Tabac</p>
        <p @click="clickSubcategory($event)" class="subcategory">Assurance</p>
        <p @click="clickSubcategory($event)" class="subcategory">Dons</p>
        <p @click="clickSubcategory($event)" class="subcategory">Divers - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/sante.png" />
        <p class="p-category-title">Santé</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Dentiste</p>
        <p @click="clickSubcategory($event)" class="subcategory">Mutuelle</p>
        <p @click="clickSubcategory($event)" class="subcategory">Médecin</p>
        <p @click="clickSubcategory($event)" class="subcategory">Pharmacie</p>
        <p @click="clickSubcategory($event)" class="subcategory">Opticien / Opthalmo</p>
        <p @click="clickSubcategory($event)" class="subcategory">Santé - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/depensePro.png" />
        <p class="p-category-title">Dépense pro</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Comptabilité</p>
        <p @click="clickSubcategory($event)" class="subcategory">Cotisation</p>
        <p @click="clickSubcategory($event)" class="subcategory">Conseil</p>
        <p @click="clickSubcategory($event)" class="subcategory">Maintenance</p>
        <p @click="clickSubcategory($event)" class="subcategory">Note de Frais</p>
        <p @click="clickSubcategory($event)" class="subcategory">
          Marketing / Publicité
        </p>
        <p @click="clickSubcategory($event)" class="subcategory">Salaire</p>
        <p @click="clickSubcategory($event)" class="subcategory">Sous traitance</p>
        <p @click="clickSubcategory($event)" class="subcategory">Dépense pro - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/esthetique.png" />
        <p class="p-category-title">Esthétique <em></em> soins</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Coiffeur</p>
        <p @click="clickSubcategory($event)" class="subcategory">Cosmétique</p>
        <p @click="clickSubcategory($event)" class="subcategory">Esthétique</p>
        <p @click="clickSubcategory($event)" class="subcategory">Spa / Massage</p>
        <p @click="clickSubcategory($event)" class="subcategory">Soins - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/taxe.png" />
        <p class="p-category-title">Impôt <em></em> taxe</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">Amende</p>
        <p @click="clickSubcategory($event)" class="subcategory">Impôt</p>
        <p @click="clickSubcategory($event)" class="subcategory">Tva</p>
        <p @click="clickSubcategory($event)" class="subcategory">Taxe</p>
        <p @click="clickSubcategory($event)" class="subcategory">Taxe - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/cheque.png" />
        <p class="p-category-title">oper° Banque</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p class="subcategory">Chèque</p>
        <p class="subcategory">Retrait</p>
        <p class="subcategory">Virement</p>
        <p class="subcategory">Oper° - Autre</p>
      </div>
    </div>
    <div class="category-P">
      <div class="category-title" @click="clickCategoryTitle($event)">
        <img class="img-category" src="@/assets/img/enfant.png" />
        <p class="p-category-title">Scolarité <em></em> enfant</p>
        <img src="@/assets/img/dropdown.png" class="dropdown" />
      </div>
      <div class="content">
        <p @click="clickSubcategory($event)" class="subcategory">
          Baby-sitters <em></em> crèche
        </p>
        <p @click="clickSubcategory($event)" class="subcategory">Ecole</p>
        <p @click="clickSubcategory($event)" class="subcategory">Fourniture scolaire</p>
        <p @click="clickSubcategory($event)" class="subcategory">Jouets</p>
        <p @click="clickSubcategory($event)" class="subcategory">Logement Etudiant</p>
        <p @click="clickSubcategory($event)" class="subcategory">Pensions</p>
        <p @click="clickSubcategory($event)" class="subcategory">prêt etudiant</p>
        <p @click="clickSubcategory($event)" class="subcategory">Scolarité - Autre</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ListCategory",
  data() {
    return {
      subcategorySelected: null,
      maincategorySelected: null
    };
  },
  methods: {
    clickCategoryTitle(event) {
      let src = event.srcElement;
      let category_P;

      if (
        src.classList.contains("p-category-title") ||
        src.classList.contains("dropdown") ||
        src.classList.contains("img-category")
      ) {
        src.parentElement.parentElement.classList.toggle("open");
        category_P = src.parentElement.parentElement;
      } else if (src.classList.contains("category-title")) {
        src.parentElement.classList.toggle("open");
        category_P = src.parentElement;
      }
      document.querySelectorAll(".category-P.open").forEach((item) => {
        if (item !== category_P) item.classList.remove("open");
      });
    },
    clickSubcategory(event) {
      const src = event.srcElement;
      if(!src.classList.contains("selected")) {
        src.classList.add("selected");
        this.subcategorySelected = src.innerHTML;
        const titleC = src.parentElement.parentElement.querySelector(".category-title .p-category-title");
        this.maincategorySelected = titleC.innerHTML;
        const subcategory = this.subcategorySelected;
        const maincategory = this.maincategorySelected;
        
        this.$emit('subcategorySelected', {subcategory, maincategory});
      }
    },
  },
};
</script>

<style scoped>
.container {
  margin: 0px auto;
  padding: 0px 0px;
  width: 250px;
  height: 300px;
  overflow: auto;
  background-color: #ffff;
  border-radius: 5px 5px;
  box-shadow: rgba(0, 0, 0, 0.25) 0px 0.0625em 0.0625em,
    rgba(0, 0, 0, 0.25) 0px 0.125em 0.5em,
    rgba(255, 255, 255, 0.1) 0px 0px 0px 1px inset;
}

.category-P {
  margin: 0px 0px;
  display: flex;
  flex-direction: column;
  padding: 0px 7px;
  margin-top: 5px;
  color: black;
}
.category-P:last-child {
  margin-bottom: 5px;
}

.category-title {
  z-index: 2;
  padding: 5px 12px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  align-content: center;
  cursor: pointer;
  border-top: solid 1px transparent;
  border-bottom: solid 1px transparent;
  transition: 0.3s;
  border-radius: 25px 25px;
  background-color: #f2f2f2;
}
.category-title p,
.category-title img {
  cursor: pointer;
}

.category-P.open .category-title {
  border-radius: 5px 5px 0px 0px;
  border-bottom: none;
}

div.category-title p {
  margin: 0px 0px;
  padding: 0px 0px;
  border-bottom: solid 1px transparent;
  transition: 0.3s;
}
div.category-title:hover p {
  border-color: black;
}

div.category-title:hover .dropdown,
.category-P.open .dropdown {
  opacity: 1;
}

.category-P.open .dropdown {
  transform: rotateZ(180deg);
}

.img-category {
  width: 25px;
  height: 25px;
}
.category-P .dropdown {
  width: 13px;
  height: 13px;
  transition: 0.3s;
  opacity: 0;
}
.category-P .content {
  height: 0px;
  overflow-y: hidden;
  transition: 0.3s;
  padding: 0px 12px;
}
.category-P.open .content {
  padding: 5px 12px;
  color: black;
  border-radius: 0px 0px 5px 5px;
  background-color: #f2f2f2;
}
.category-P:nth-child(1).open .content {
  height: 220px;
}
.category-P:nth-child(2).open .content {
  height: 140px;
}
.category-P:nth-child(3).open .content {
  height: 280px;
}
.category-P:nth-child(4).open .content {
  height: 230px;
}
.category-P:nth-child(5).open .content {
  height: 170px;
}
.category-P:nth-child(6).open .content {
  height: 250px;
}
.category-P:nth-child(7).open .content {
  height: 90px;
}
.category-P:nth-child(8).open .content {
  height: 90px;
}
.category-P:nth-child(9).open .content {
  height: 170px;
}
.category-P:nth-child(10).open .content {
  height: 250px;
}
.category-P:nth-child(11).open .content,
.category-P:nth-child(12).open .content {
  height: 140px;
}
.category-P:nth-child(13).open .content {
  height: 120px;
}
.category-P:nth-child(14).open .content {
  height: 230px;
}

.category-P .content p {
  margin: 0px 0px;
  padding: 4px 0px;
  opacity: 0;
  transition: 0.3s;
  color: black;
  border-radius: 5px 5px;
  text-align: left;
  padding-left: 10px;
  cursor: pointer;
}

.category-P.open .content p {
  opacity: 1;
}
.category-P.open .content p:hover {
  background-color: #e8e8e8;
}
</style>
