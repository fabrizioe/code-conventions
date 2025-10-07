module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    // Allow longer commit message body (default is 100, we'll set to 600)
    'body-max-length': [2, 'always', 600],
    
    // Allow longer commit message subject (increase from 72 to 100)
    'subject-max-length': [2, 'always', 100],
    
    // Ensure body has max line length of 250 chars (increased to allow longer lines)
    'body-max-line-length': [2, 'always', 250],
    
    // Keep other conventional commit rules
    'type-enum': [
      2,
      'always',
      [
        'feat',
        'fix',
        'docs',
        'style',
        'refactor',
        'perf',
        'test',
        'build',
        'ci',
        'chore'
      ]
    ],
    'type-case': [2, 'always', 'lower-case'],
    'type-empty': [2, 'never'],
    'scope-case': [2, 'always', 'lower-case'],
    'subject-empty': [2, 'never'],
    'subject-full-stop': [2, 'never', '.'],
    'header-max-length': [2, 'always', 100]
  }
};