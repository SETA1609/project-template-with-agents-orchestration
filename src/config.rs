use crate::transport::TransportKind;
use clap::Parser;
use std::env;

/// Runtime configuration resolved at startup.
/// ENV vars are loaded first; CLI flags win if both are set.
#[derive(Debug)]
pub struct Config {
    pub transport: TransportKind,
    pub host: String,
    pub port: u16,
    pub log_level: String,
}

#[derive(Parser)]
struct Args {
    /// Transport mode: stdio or http
    #[arg(long)]
    transport: Option<String>,
}

impl Config {
    pub fn load() -> Self {
        // Load .env file if present
        dotenvy::dotenv().ok();

        let args = Args::parse();

        let transport_str = args.transport
            .or_else(|| env::var("MCP_TRANSPORT").ok())
            .unwrap_or_else(|| "http".to_string());

        let transport = match transport_str.as_str() {
            "stdio" => TransportKind::Stdio,
            "http" => TransportKind::Http,
            _ => TransportKind::Http,
        };

        let host = env::var("HOST").unwrap_or_else(|_| "0.0.0.0".to_string());
        let port = env::var("PORT")
            .ok()
            .and_then(|p| p.parse().ok())
            .unwrap_or(3000);
        let log_level = env::var("LOG_LEVEL").unwrap_or_else(|_| "info".to_string());

        Self {
            transport,
            host,
            port,
            log_level,
        }
    }
}