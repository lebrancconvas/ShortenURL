<script lang="ts">
  import axios from 'axios';
  let url: string = '';
  let shortURL: string = '';
  let isCopy: boolean = false;
  let handleSubmit = async() => {
    axios.post('http://localhost:8101/api/v1/short/url', {
      "original_url": url
    }).then(res => {
      shortURL = res.data.short_url;
      console.log(res.data.short_url);
    }).catch(err => {
      console.error(err);
    })
  };

  let openURL = () => {
    console.log(`Open: ${shortURL}`);
    window.open(`http://localhost:8101/${shortURL}`, '_blank');
  }

  let copyURL = () => {
    console.log(`Copy: ${shortURL}`);
    isCopy = true;
    navigator.clipboard.writeText(`http://localhost:8101/${shortURL}`);
  }
</script>

<svelte:head>
  <title>URL Shortner</title>
</svelte:head>

<main>
  <header>
    <div>
      <h1>URL Shortener</h1>
    </div>
    <div>
      <h3>(Recommended to paste the url in http or https format)</h3>
    </div>
  </header>
  <section id="form-section">
    <form on:submit={handleSubmit}>
      <input type="text" bind:value={url} id="url" placeholder="Enter URL">
      <input type="submit" id="submit" value="Shorten">
    </form>
  </section>
  {#if shortURL !== ''}
  <section id="short-url-action-section">
    <div>
      <h2>Short URL: http://localhost:8101/{shortURL}</h2>
      <div id="action">
        <button id="open-url" on:click={openURL}>Open URL</button>
        {#if !isCopy}
        <button id="copy-url" on:click={copyURL}>Copy URL</button>
        {:else}
        <button id="copy-url" disabled>Copied</button>
        {/if}
      </div>
    </div>
  </section>
  {/if}
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

  #short-url-action-section {
    margin-top: 50px;
  }

  #open-url {
    border: none;
    border-radius: 5px;
    background-color: rgb(229, 147, 47);
    color: white;
    padding: 10px;
    font-size: 20px;
    cursor: pointer;
  }

  #copy-url {
    border: none;
    border-radius: 5px;
    background-color: rgb(24, 186, 92);
    color: white;
    padding: 10px;
    font-size: 20px;
    cursor: pointer;
  }

  #open-url:active, #copy-url:active {
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
