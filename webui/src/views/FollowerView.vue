<script>
export default{
  data(){
    return{
      ID : localStorage.getItem('IDCercato'),
      seguaci: []
    }
  },
  methods:{
    async refresh(){
      this.recuperaInfo();
    },
    async recuperaInfo(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + this.ID, {headers: {Authorization: "Bearer " + this.ID}});
        
        if (response.data.Followers != null) {this.seguaci = response.data.Followers; }
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
  },
  mounted (){
    this.refresh();
  }
}
</script>

<template>
  <body>
    <header class=intestazione-seguaci>
      <h2 style = "font-weight: bold; padding-left: 20px; cursor:pointer;"
      @click = "() => {this.$router.back();}"> Torna indietro </h2>
      <h1 style = "font-weight: bold;">Account Seguaci</h1>
      <h2 style = "padding-right: 20px;"> Seguaci : {{this.seguaci.length}}</h2>
    </header>

    <section class = lista-seguaci>
      <div class = follower v-for = "utente in this.seguaci" v-if = "this.seguaci.length > 0">
        <h2>{{utente.Username}}</h2>
        <button class = blocca style = "color:red">Blocca</button>
      </div>
    </section>
  </body>
</template>

<style>
  body{
    margin: 0;
    margin-right: 0;
    width:100%;

    display: inline-block;
    background-color: whitesmoke;
  }

  .intestazione-seguaci{
    position: fixed;
    width:100%;
    height:100px;

    display: flex;
    align-items:center;
    justify-content: space-between;

    background-color: cadetblue;
    border-bottom: 5px solid black;
  }

  .lista-seguaci{
    padding-left: 10%;
    padding-right: 10%;
    padding-top:120px;
    width: 100%;
    
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 10px;
  }

  .follower{
    cursor: pointer;
    border-radius: 10px;
    box-shadow: 10px 10px 5px cadetblue;
    text-align: center;
    border: 2px solid black;
  }

</style>