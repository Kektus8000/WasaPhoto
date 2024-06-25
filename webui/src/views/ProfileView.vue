<script>
export default{
  data(){
    return{
      errormsg: "",
      profilo : {
        ID : localStorage.getItem('IDCercato'),
        nome : "",
        seguaci : [],
        seguiti : [],
        bannati : [],
        fotoPubblicate : []
      }
    }
  },
  methods: {
    async refresh(){
      await this.recuperaInfo();
    },
    async recuperaInfo(){
      try
      {
        let response = await this.$axios.get('/userProfile/' + localStorage.getItem('IDCercato'), {headers: {Authorization: "Bearer " + localStorage.getItem('identifier')}});
        this.profilo.nome = response.data.Username;
        if (response.data.Followers != null) {for (let i = 0; i < response.data.Followers.length; i++){seguaci.push(response.data.Followers[i]);} }
        if (response.data.Followings != null) {for (let i = 0; i < response.data.Followings.length; i++){seguiti.push(response.data.Followings[i]);} }
        if (response.data.Banneds != null) {for (let i = 0; i < response.data.Banneds.length; i++){bannati.push(response.data.Banneds[i]);} }
        if (response.data.PublishedPhotos != null) {for (let i = 0; i < response.data.PublishedPhotos.length; i++){fotoPubblicate.push(response.data.PublishedPhotos[i]);} }
      }
      catch(e)
      {
        this.errormsg = e.toString();
        alert(this.errormsg);
      }
    },
    async controllaSeguiti(){
      this.$router.push({path: '/userProfile/${this.profilo.ID}/following'})
    },
    async tornaHomePage(){
      localStorage.removeItem('IDCercato');
      this.$router.replace('/session');
    }
  },
  mounted(){
    this.recuperaInfo();
  }
}
</script>

<template>
  <body>
    <header class=intestazione>
      <h4 @click = "tornaHomePage" >Torna alla HomePage</h4>
      <h1 style = "font-weight: bold;">Profilo di {{this.profilo.nome}}</h1>
      <div class=statistiche>
        <h4 @click = "controllaSeguiti">Followers: {{this.profilo.seguaci.length}}</h4>
        <h4>Seguiti: {{this.profilo.seguiti.length}}</h4>
      </div>
    </header>

    <div class = riga>
      <template class = listaFollowers v-for = "utente in profilo.fotoPubblicate">
        <div class = colonna>
          <img class = foto :src = utente.link alt="foto">
        </div>
      </template>
    </div>
  </body>
</template>

<style>
  body{
    margin: 0;
    margin-right: 0;
    width:100%;

    background-color: whitesmoke;
  }

  .intestazione{
    width:100%;
    display: flex;
    align-items: center;
    justify-content:space-between;
    cursor: pointer;
    height:100px;

    background-color:brown;
    text-align: left;

    border-bottom-left-radius: 30px;
    border-bottom-right-radius: 30px;
    border-bottom: 3px solid black;
  }

  .riga{
    display: flex;
    flex-wrap: wrap;
  }

  .colonna{
    flex: 1 0 30%;
    padding: 5px;
  }

  .foto{
    width: 100%;
    height:100%;
    border-radius: 10px;;
  }
</style>