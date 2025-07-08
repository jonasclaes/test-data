# Test Data Repository

A collection of sample files for automated testing, quality assurance, and development workflows.

## 🎯 Purpose

This repository provides a variety of test files including documents, images, media files, and data formats that can be used for:

- **Automated testing scripts** - File processing and validation tests
- **Quality assurance workflows** - Consistent test scenarios across environments  
- **CI/CD pipelines** - Reliable test data for continuous integration
- **Development and debugging** - Sample files for local testing

## 🚀 Use

To use files from this repository in your tests or applications, use the GitHub raw content endpoint:

```
https://raw.githubusercontent.com/jonasclaes/test-data/main/path/to/your/file.ext
```

For version-specific access, use a commit hash instead of the branch name:

```
https://raw.githubusercontent.com/jonasclaes/test-data/{commit-hash}/path/to/your/file.ext
```

**Example:**
```bash
# Download a test file directly (latest version)
curl -O https://raw.githubusercontent.com/jonasclaes/test-data/main/pdfs/basic.pdf

# Download from a specific commit
curl -O https://raw.githubusercontent.com/jonasclaes/test-data/abc123def456/pdfs/basic.pdf
```

**Benefits of using permalinks:**
- **Reliable URLs** - Files are accessible from any environment
- **Version control** - Use specific commits for consistent test data
- **Immutable references** - Commit hashes ensure the exact same file content
- **No local storage** - Fetch files on-demand in CI/CD pipelines
- **Easy integration** - Works with any HTTP client or testing framework

## 🤝 Contributing

Contributions are welcome! Feel free to add new test files that would be useful for the community:

1. Fork the repository
2. Add your test files to appropriate directories
3. Submit a pull request

**Guidelines:**
- Avoid files containing sensitive or personal information
- Keep file sizes reasonable (under 10MB when possible)
- Use descriptive file names

## ⚖️ License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

You are free to use, modify, and distribute these test files for any purpose, including commercial use.

---

**Note**: This repository is intended for testing purposes. Files should not contain real personal data or sensitive information.
