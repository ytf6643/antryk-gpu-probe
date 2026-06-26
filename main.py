import subprocess

from fastapi import FastAPI


app = FastAPI()


def run_command(args):
    try:
        return subprocess.check_output(
            args,
            text=True,
            stderr=subprocess.STDOUT,
            timeout=15,
        )
    except Exception as exc:
        return f"{type(exc).__name__}: {exc}"


@app.get("/")
def root():
    return {
        "status": "ok",
        "nvidia_smi": run_command(["nvidia-smi"]),
    }


@app.get("/health")
def health():
    return {"status": "ok"}
