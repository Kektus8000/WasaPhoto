<template>
  <div class = "Titolo">
      <h1 class="Profilo"> Profilo di {{this.username}} </h1>
  </div>

  <nav class ="BarraLaterale">
    <ul class = "listaOpzioni">
      <li class="home">
        <a>
          <h3>Home Page</h3>
        </a>
      </li>
      <li>
        <a>
          <h3>Seguiti : {{this.seguiti}}</h3>
        </a>
      </li>
      <li>
        <a>
          <h3>Followers : {{this.followers}}</h3>
        </a>
      </li>
    </ul>
  </nav>

  <main class = "SchermataPrincipale">
    <div class = "FotoPubblicate">
      <ul v-for="foto in immagini">
        <li>
            <img class= "Foto" :src = foto.link>
        </li>
      </ul>
    </div>

  </main>
</template>

<script>
export default{
  data(){
    return{
      seguiti: -1,
      followers: this.ottieniProfilo(),
      ID: localStorage.getItem('identifier'),
      username: localStorage.getItem('username'),
      immagini: [{link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'}]
    }
  },
  methods:{
    async ottieniProfilo(){
      let response = await this.$axios.get("/userProfile/" + this.ID,
        { headers: {'Authorization': 'Bearer ' + this.ID}});
    }
  }
}
</script>

<style>
  .BarraLaterale{
    position:fixed;
    padding-top: 10px;
    top:0;
    background-color:chocolate;
    height:100%;
    width:160px;
    border-right:1px solid black;
    overflow-x: hidden;
  }

  .Titolo{
    background-color:chocolate;
    text-align: center;
    width:100%;
    border-bottom: 1px solid black;
    padding-left: 100px;
    top:0;
    position: absolute;
  }

  .Foto{
    cursor: pointer;
    box-shadow: 10px 10px;
    width: 260px;
    height: 260px;
  }

  .FotoPubblicate{
    padding-top: 100px;
    padding-left: 200px;
    display: flex;
    flex-wrap: wrap;
    flex: 0 0 33.33333%;
    gap: 10px;
    margin: 0px;
  }

  .listaOpzioni{
    padding-left: 10px;
    padding-top: 50px;
  }

  ul {
    list-style: none;
  }

</style>