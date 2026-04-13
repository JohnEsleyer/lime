<script>
    export let onOpenWorkspace;
</script>

<div class="project-manager">
    <div class="hero">
        <h1>🍋 LIME</h1>
        <p>Professional Programmatic Video Editor</p>
    </div>

    <div class="actions">
        <button class="btn btn-primary" on:click={onOpenWorkspace}>
            Open Project
        </button>
        <button class="btn btn-secondary" disabled>
            New Project (Coming Soon)
        </button>
    </div>

    <div class="recent-projects">
        <h3>Recent Projects</h3>
        <div class="project-list">
            <div class="project-item empty">No recent projects found.</div>
        </div>
    </div>
</div>

<style>
    .project-manager {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100vh;
        background: #000000;
        color: white;
        font-family: "Inter", sans-serif;
    }

    .hero {
        text-align: center;
        margin-bottom: 40px;
    }

    .hero h1 {
        font-size: 5rem;
        margin: 0;
        background: linear-gradient(135deg, #a3e635, #65a30d);
        -webkit-background-clip: text;
        background-clip: text;
        -webkit-text-fill-color: transparent;
        font-weight: 800;
    }

    .hero p {
        color: #a3a3a3;
        font-size: 1.2rem;
    }

    .actions {
        display: flex;
        gap: 20px;
        margin-bottom: 60px;
    }

    .btn {
        padding: 12px 32px;
        border-radius: 8px;
        font-weight: 600;
        cursor: pointer;
        transition: all 0.2s;
        border: none;
        font-size: 1rem;
    }

    .btn-primary {
        background: #a3e635;
        color: #0f172a;
    }

    .btn-primary:hover {
        background: #bef264;
        transform: translateY(-2px);
    }

    .btn-secondary {
        background: rgba(255, 255, 255, 0.05);
        color: white;
        border: 1px solid rgba(255, 255, 255, 0.1);
    }

    .btn-secondary:not(:disabled):hover {
        background: rgba(255, 255, 255, 0.1);
        transform: translateY(-2px);
    }

    .btn:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }

    .recent-projects {
        width: 100%;
        max-width: 600px;
    }

    .recent-projects h3 {
        color: #64748b;
        font-size: 0.9rem;
        text-transform: uppercase;
        letter-spacing: 0.1em;
        margin-bottom: 16px;
        border-bottom: 1px solid rgba(255, 255, 255, 0.05);
        padding-bottom: 8px;
    }

    .project-list {
        display: flex;
        flex-direction: column;
        gap: 12px;
    }

    .project-item.empty {
        text-align: center;
        padding: 40px;
        background: rgba(255, 255, 255, 0.02);
        border: 1px dashed rgba(255, 255, 255, 0.1);
        border-radius: 8px;
        color: #475569;
    }
</style>
