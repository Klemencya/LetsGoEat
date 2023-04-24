<script lang="ts">
    import {base} from "$app/paths";
</script>

<meta property="og:title" content="Let's Go Eat"/>
<meta property="og:description"
      content="This app has been specially designed for people who can't find company to eat."/>


<div class="form">
    <h1>Let's Go Eat</h1>
    <div id="info-block">
        <p>This app has been specially designed for people who can't find company to eat.</p>
        <p>Just choose your company and eat</p>
    </div>
    <nav>
        <a href="{base}/login">
            <button>Log in</button>
        </a>
        <a href="{base}/registration">
            <button>Registration</button>
        </a>
    </nav>
</div>

<style>
    @import url("https://fonts.googleapis.com/css2?family=Shrikhand&family=Ubuntu:wght@400;700&display=swap");

    .form {
        text-align: center;
        margin-right: 30%;
        margin-left: 30%;
        margin-top: 10%;
    }

    div h1 {
        font-size: 64px;
        font-family: "Shrikhand", cursive;
        color: #9e4eca;
    }

    div nav {
        display: flex;
        flex-direction: column;
    }

    nav button {
        margin-top: 10px;
        width: 50%;
    }

    button {
        width: 200px;
        font-size: 20px;
        font-family: "Ubuntu", sans-serif;
        font-weight: 700;
        color: #fff;
        cursor: pointer;
        margin: 20px;
        height: 55px;
        text-align: center;
        border: none;
        background-size: 300% 100%;

        border-radius: 50px;
        moz-transition: all 0.4s ease-in-out;
        -o-transition: all 0.4s ease-in-out;
        -webkit-transition: all 0.4s ease-in-out;
        transition: all 0.4s ease-in-out;
        background-image: linear-gradient(
                to right,
                #9e4eca,
                #7584b8,
                #8abcc3,
                #85d2ac
        );
        box-shadow: 0 4px 15px 0 rgba(158, 184, 198, 0.75);
    }

    button:hover {
        background-position: 100% 0;
        moz-transition: all 0.4s ease-in-out;
        -o-transition: all 0.4s ease-in-out;
        -webkit-transition: all 0.4s ease-in-out;
        transition: all 0.4s ease-in-out;
    }

    #info-block {
        padding-bottom: 50px;
    }
</style>
