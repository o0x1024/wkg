<template>
    <div class="tag-main">
        <div class="tag-container">
            <div class="tag-post-wrap tags">
                <h2 class="tag-post-title">-&nbsp;Categories&nbsp;-</h2>
                <div class="tag-cloud-tags">
                    <span class="desc">二进制、web渗透、HIDS...</span>

                </div>
            </div>
        </div>
    </div>
</template>


<style>
.tag-main {
    flex-grow: 1;
    flex-shrink: 0;
    flex-basis: auto;
    display: flex;
    flex-direction: column;
}


.tag-main .tag-container {
    padding-left: 1em;
    padding-right: 1em;
    height: 100%;
    display: flex;
    flex-direction: column;
    flex: 1;
}

.tag-post-wrap {
    position: relative;
    width: 100%;
    max-width: 780px;
    margin: 0 auto;
    padding-top: 2rem;
}

.tag-post-wrap .tag-post-title {
    font-size: 2em;
    line-height: 1.5em;
}

.tag-cloud-tags {
    margin: 10px 0;
    padding-top: 2em;
}

.tag-cloud-tags a {
    display: inline-block;
    position: relative;
    margin: 5px 10px;
    word-wrap: break-word;
    transition-duration: 0.3s;
    transition-property: transform;
    transition-timing-function: ease-out;
}

a {
    color: #161209;
    text-decoration: none;
    transition: color 0.2s ease, border-color 0.2s ease, background 0.2s ease, opacity 0.2s ease;
    cursor: pointer;
}
</style>