<script>
export default{
  data(){
    return{
      errormsg: "",
      profilo: {
        ID : localStorage.getItem('profiloRicercato').ID,
        nome : localStorage.getItem('profiloRicercato').nome,
        seguiti : localStorage.getItem('profiloRicercato').seguiti,
        seguaci : localStorage.getItem('profiloRicercato').seguaci,
        bannati : localStorage.getItem('profiloRicercato').bannati,
        fotoPubblicate : localStorage.getItem('profiloRicercato').fotoPubblicate
      },
      identifier: localStorage.getItem('identifier'),
      username: localStorage.getItem('username'),
      seguiti: [{link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                {link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                {link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'},
                {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                {link: 'https://upload.wikimedia.org/wikipedia/it/2/22/Dragon_Ball_Super.png'},
                {link:'https://www.avvenire.it/c/2017/PublishingImages/debda429421d455a8975d6ba03e67d65/caparezza.jpg?width=1024'},
                {link: 'https://cdn-2.motorsport.com/images/amp/0ZRKlvo0/s1000/formula-1-spanish-gp-2023-char-2.jpg'}]
    }
  }
}
</script>

<template>
  <body>
    <header class=intestazione>
      <RouterLink :to = "'/session'">
        <h4 class=home>Torna alla HomePage</h4>
      </RouterLink>
      <img class = icona src = "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2b/Alonso_2016.jpg/640px-Alonso_2016.jpg">  
      <h1 class = nome>Profilo di {{this.profilo.nome}}</h1>
      <div class=statistiche>
        <h4>Followers: {{this.profilo.seguaci.length}}</h4>
        <h4>Seguiti: {{this.profilo.seguiti.length}}</h4>
      </div>
    </header>

    <div class = riga>
      <template class = listaFollowers v-for = "utente in this.seguiti">
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

  .nome{
    float: left;
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

  .icona{
    border-radius: 50%;
    border-style: solid;
    border-color: rgba(0,0,0,0.5);

    width:50px;
  }

</style>