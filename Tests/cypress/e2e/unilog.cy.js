describe('toolshop', () => {
  it('open site', () => {
    cy.visit('https://practicesoftwaretesting.com/')
  })

  it('nav: categories', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="nav-categories"]').click()
    cy.get('.dropdown-menu').should('be.visible')
  })

  it('nav: category hand tools', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="nav-categories"]').click()
    cy.get('[data-test="nav-hand-tools"]').click()

    cy.get('[data-test="page-title"]').contains("Hand Tools")
  })

  it('nav: contact', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="nav-contact"]').click({ force: true })
    cy.location('pathname').should('eq', '/contact')
  })

  it('nav: sign in', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="nav-sign-in"]').click({ force: true })
    cy.location('pathname').should('eq', '/auth/login')
  })

  it('nav: lang DE', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="language-select"]').click()
    cy.get('[data-test="lang-de"]').click()

    cy.get('.container-fluid').contains('Das ist eine Demo Applikation')
  })

  it('nav: lang ES', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="language-select"]').click()
    cy.get('[data-test="lang-es"]').click()

    cy.get('.container-fluid').contains('Esta es una aplicación DEMO')
  })

  it('nav: lang FR', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="language-select"]').click()
    cy.get('[data-test="lang-fr"]').click()

    cy.get('.container-fluid').contains('Ceci est une application de démonstration')
  })

  it('nav: lang NL', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="language-select"]').click()
    cy.get('[data-test="lang-nl"]').click()

    cy.get('.container-fluid').contains('Dit is een DEMO-applicatie')
  })

  it('nav: lang TR', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="language-select"]').click()
    cy.get('[data-test="lang-tr"]').click()

    cy.get('.container-fluid').contains('Bu bir DEMO uygulamasıdır')
  })

  it('privacy policy', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[routerlink="privacy"]').click()

    cy.get('body').contains('Privacy Policy for Toolshop')
  })


  it('contact: empty form error', () => {
    cy.visit('https://practicesoftwaretesting.com/contact')

    cy.get('[data-test="contact-submit"]').click()

    cy.get('[data-test="first-name-error"] > div').should('be.visible')
    cy.get('[data-test="last-name-error"] > div').should('be.visible')
    cy.get('[data-test="email-error"] > div').should('be.visible')
    cy.get('[data-test="subject-error"] > div').should('be.visible')
    cy.get('[data-test="message-error"] > div').should('be.visible')
  })

  it('contact form', () => {
    cy.visit('https://practicesoftwaretesting.com/contact')

    cy.get('[data-test="first-name"]').type('Test')
    cy.get('[data-test="last-name"]').type('Test')
    cy.get('[data-test="email"]').type('test@test.com')

    cy.get('[data-test="subject"]').select('Webmaster')
    cy.get('[data-test="message"]').type('Test test test test test test test test test test test test test test')

    cy.get('[data-test="contact-submit"]').click()

    cy.get('.alert').should('be.visible')
    cy.get('.alert').contains('Thanks for your message! We will contact you shortly.')
  })

  it('sign in: empty form error', () => {
    cy.visit('https://practicesoftwaretesting.com/auth/login')

    cy.get('[data-test="login-submit"]').click()

    cy.get('[data-test="email-error"]').should('be.visible')
    cy.get('[data-test="password-error"]').should('be.visible')
  })

  it('forgot password: empty form error', () => {
    cy.visit('https://practicesoftwaretesting.com/auth/forgot-password')

    cy.get('[data-test="forgot-password-submit"]').click()
    cy.get('[data-test="email-error"]').should('be.visible')
  })

  it('register: empty form error', () => {
    cy.visit('https://practicesoftwaretesting.com/auth/register')

    cy.get('[data-test="register-submit"]').click()

    cy.get('[data-test="first-name-error"]').should('be.visible')
    cy.get('[data-test="email-error"]').should('be.visible')
    cy.get('[data-test="dob-error"]').should('be.visible')

    cy.get('[data-test="street-error"]').should('be.visible')
    cy.get('[data-test="postal_code-error"]').should('be.visible')
    cy.get('[data-test="city-error"]').should('be.visible')
    cy.get('[data-test="state-error"]').should('be.visible')
    cy.get('[data-test="country-error"]').should('be.visible')

    cy.get('[data-test="phone-error"]').should('be.visible')
    cy.get('[data-test="email-error"]').should('be.visible')
    cy.get('[data-test="password-error"]').should('be.visible')
  })

  it('search', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="search-query"]').type('Thor')
    cy.get('[data-test="search-submit"]').click()
    
    cy.get('[data-test="search-caption"]').contains('Thor')
    cy.get('[data-test="product-name"]').should('be.visible')
    cy.get('[data-test="product-name"]').contains('Thor Hammer')
  })

  it('search lowcase', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="search-query"]').type('thor')
    cy.get('[data-test="search-submit"]').click()
    
    cy.get('[data-test="search-caption"]').contains('thor')
    cy.get('[data-test="product-name"]').should('be.visible')
    cy.get('[data-test="product-name"]').contains('Thor Hammer')
  })

  it('search uppercase', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="search-query"]').type('THOR')
    cy.get('[data-test="search-submit"]').click()
    
    cy.get('[data-test="search-caption"]').contains('THOR')
    cy.get('[data-test="product-name"]').contains('Thor Hammer')
  })

  it('search with spaces', () => {
    cy.visit('https://practicesoftwaretesting.com/')

    cy.get('[data-test="search-query"]').type('       thor    ')
    cy.get('[data-test="search-submit"]').click()
    
    cy.get('[data-test="search-caption"]').contains('thor')
    cy.get('[data-test="product-name"]').contains('Thor Hammer')
  })
})
