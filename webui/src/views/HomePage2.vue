<script>
export default{
  data(){
    return{
      ricerca: "",
      username: "Simone",
      immagini: [{link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                  {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                  {link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                  {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'}]
    }
  },
  methods:{
    stampa(){console.log("Simone");}
  }
}
</script>

<template>
  <body>
    <section>
      <header class="intestazione">
        <h1 class = "benvenuto">Bentornato {{this.username}}</h1>
      </header>
    </section>

    <br>
    <section>
      <ul class = fotoDisplayer v-for = "foto in this.immagini">
        <div class = fotoContainer>
          <img class = immagine :src= foto.link @click=stampa>
          <button class = like>Mi piace</button>
          <button class = commento>Commenti</button>
        </div>  
      </ul>
    </section>
  </body>
</template>

<style>
  .intestazione{
    position:absolute;
    flex: 1;
    width:100%;
    top: 0%;
    left:0%;
    background-color: orange;
    border-bottom-left-radius: 10px;
    border-bottom-right-radius: 10px;

    font-size:20px;
    text-align: center;
    color:white;  
    font-family:'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    -webkit-text-stroke-width: 1px;
    -webkit-text-stroke-color: black;
  
  }
  
  .fotoContainer{
    border: 1px;
    border-color: black;

    padding-top: 40px;
  }

  .immagine {
    cursor: pointer;
    width: 100%;
  }

  .like{
    position: relative;
    width: 30%;
    transform: translate(20%, -140%);

    text-align:center;
    border-radius: 1px;
    font-family:'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    cursor: pointer;
  }

  .commento{
    position: relative;
    width: 30%;
    transform: translate(110%, -140%);

    text-align:center;
    border-radius: 1px;
    font-family:'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    cursor: pointer;
  }
</style>

