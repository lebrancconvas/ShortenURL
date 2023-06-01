<script lang="ts">
  import axios from 'axios';
  let url: string = '';
  let handleSubmit = async() => {
    axios.post('http://localhost:8101/api/v1/short/url', {
      "original_url": url
    }).then(res => {
      console.log(res.data);
    }).catch(err => {
      console.error(err);
    })
  };
</script>

<svelte:head>
  <title>URL Shortner</title>
</svelte:head>

<main>
  <header>
    <div>
      <h1>URL Shortener</h1>
    </div>
  </header>
  <section id="form-section">
    <form on:submit={handleSubmit}>
      <input type="text" bind:value={url} id="url" placeholder="Enter URL">
      <input type="submit" id="submit" value="Shorten">
    </form>
  </section>
  <footer>
    <div>
      <p>Copyright 2023 <a href="https://github.com/lebrancconvas" target="_blank">Poom Yimyuean (@lebrancconvas)</a></p>
    </div>
  </footer>
</main>

<style>
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Chakra Petch', sans-serif;
    text-align: center;
  }

  h1 {
    font-size: 45px;
  }

  #url {
    width: 30%;
    height: 30px;
    font-size: 20px;
  }

  #submit {
    border: none;
    border-radius: 5px;
    background-color: rgb(49, 148, 241);
    color: white;
    padding: 10px;
    font-size: 20px;
  }

  #submit:hover {
    cursor: pointer;
    background-color: rgb(49, 148, 241, 0.8);
  }

  #submit:active {
    transform: scale(0.95);
  }


  footer {
    position: absolute;
    bottom: 0;
    margin: 0;
    padding: 0;
    width: 95%;
    height: 50px;
    background-color: white;
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>
