
const myCodeMirrorIn = CodeMirror.fromTextArea(document.querySelector('#soql'), {mode:'sql', theme: 'dracula', lineNumbers: true});
let lastDiagram = '';

function execVisualize(event) {
    event?.preventDefault();
    try {
        let result = visualizeSoql(myCodeMirrorIn.getDoc().getValue());

        const container = document.getElementById('erDiagramContainer');
        container.removeAttribute('data-processed');

        lastDiagram = result;

        container.className = 'mermaid';
        container.textContent = result;
        mermaid.init(undefined, container);
    } catch(e) {
        rebootGoApplication();
    }
}

async function copyDiagramDefToClipboard(event) {
    event.preventDefault();
    try {
        await navigator.clipboard.writeText(lastDiagram);
        const el = document.querySelector('#copied');
        el.style.display = 'inline-block';
        setTimeout(() => {
            el.style.display = 'none';
        }, 1200);
    } catch (e) {
        //
    }
}

async function copySvgToClipboard(event) {
    event.preventDefault();
    const container = document.getElementById('erDiagramContainer');
    const svg = container.querySelector('svg');
    const text = svg.outerHTML;
    try {
        await navigator.clipboard.writeText(text);
        const el = document.querySelector('#copied');
        el.style.display = 'inline-block';
        setTimeout(() => {
            el.style.display = 'none';
        }, 1200);
    } catch (e) {
        //
    }
}
