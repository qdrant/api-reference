export function Footer() {
  return (
    <footer className="main-footer">
      <div className="container-xxl">
        <div className="widgets-section">
          <div className="row clearfix">
            <div className="main-footer__column col-12">
              <div className="main-footer__widget">
                <div className="main-footer__title">
                  <h4>Product</h4>
                </div>
                <ul className="main-footer__list">
                  <li>
                    <a href="https://qdrant.tech/use-cases/">
                      <span>Use cases</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.tech/solutions/">
                      <span>Solutions</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.tech/benchmarks/">
                      <span>Benchmarks</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.tech/demo/">
                      <span>Demos</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.tech/pricing/">
                      <span>Pricing</span>
                    </a>
                  </li>
                </ul>
              </div>
              <div className="main-footer__widget">
                <div className="main-footer__title">
                  <h4>Community</h4>
                </div>
                <ul className="main-footer__list">
                  <li>
                    <a href="https://github.com/qdrant/qdrant" target="_blank">
                      <i className="fab fa-github"></i>&nbsp;
                      <span>Github</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.to/discord" target="_blank">
                      <i className="fab fa-discord"></i>&nbsp;
                      <span>Discord</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.to/twitter" target="_blank">
                      <i className="fab fa-twitter"></i>&nbsp;
                      <span>Twitter</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.tech/subscribe/">
                      <i className="fas fa-mail-bulk"></i>&nbsp;
                      <span>Newsletter</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.to/contact-us" target="_blank">
                      <i className="fas fa-envelope"></i>&nbsp;
                      <span>Contact us</span>
                    </a>
                  </li>
                </ul>
              </div>
              <div className="main-footer__widget">
                <div className="main-footer__title">
                  <h4>Company</h4>
                </div>
                <ul className="main-footer__list">
                  <li>
                    <a href="https://qdrant.join.com">
                      <span>Jobs</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.tech/legal/privacy-policy/">
                      <span>Privacy Policy</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.tech/legal/terms_and_conditions/">
                      <span>Terms</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.tech/legal/impressum/">
                      <span>Impressum</span>
                    </a>
                  </li>
                  <li>
                    <a href="https://qdrant.tech/legal/credits/">
                      <span>Credits</span>
                    </a>
                  </li>
                </ul>
              </div>
              <div className="main-footer__widget news-widget">
                <div className="main-footer__title">
                  <h4>Latest Publications</h4>
                </div>
                <div className="widget-content">
                  <div className="post">
                    <div className="thumb">
                      <a href="https://qdrant.tech/articles/rag-is-dead/">
                        <img
                          src="https://qdrant.tech/articles_data/rag-is-dead/icon.svg"
                          alt=""
                        />
                      </a>
                    </div>
                    <h4>
                      <a href="https://qdrant.tech/articles/rag-is-dead/">
                        Why are vector databases needed for RAG? We debunk
                        claims of increased LLM accuracy and look into drawbacks
                        of large context windows.
                      </a>
                    </h4>
                  </div>
                  <div className="post">
                    <div className="thumb">
                      <a href="https://qdrant.tech/articles/binary-quantization-openai/">
                        <img
                          src="https://qdrant.tech/articles_data/binary-quantization-openai/icon.svg"
                          alt=""
                        />
                      </a>
                    </div>
                    <h4>
                      <a href="https://qdrant.tech/articles/binary-quantization-openai/">
                        Use Qdrant's Binary Quantization to enhance OpenAI
                        embeddings
                      </a>
                    </h4>
                  </div>
                  <div className="post">
                    <div className="thumb">
                      <a href="https://qdrant.tech/articles/multitenancy/">
                        <img
                          src="https://qdrant.tech/articles_data/multitenancy/icon.svg"
                          alt=""
                        />
                      </a>
                    </div>
                    <h4>
                      <a href="https://qdrant.tech/articles/multitenancy/">
                        Combining our most popular features to support scalable
                        machine learning solutions.
                      </a>
                    </h4>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        {/* <ul className="social-box">
          <li>
            <a href="https://github.com/qdrant/qdrant" target="_blank">
              <i className="fab fa-github"></i>
            </a>
          </li>
          <li>
            <a href="https://qdrant.to/linkedin" target="_blank">
              <i className="fab fa-linkedin"></i>
            </a>
          </li>
          <li>
            <a href="https://qdrant.to/twitter" target="_blank">
              <i className="fab fa-twitter"></i>
            </a>
          </li>
          <li>
            <a href="https://qdrant.to/discord" target="_blank">
              <i className="fab fa-discord"></i>
            </a>
          </li>
          <li>
            <a
              href="https://www.youtube.com/channel/UC6ftm8PwH1RU_LM1jwG0LQA"
              target="_blank">
              <i className="fab fa-youtube"></i>
            </a>
          </li>
        </ul> */}
      </div>
      <div className="footer-bottom">
        <div className="auto-container">
          <div className="row clearfix">
            <div className="copyright-column col-lg-6 col-md-12 col-sm-12">
              <div className="copyright">
                © 2024 Qdrant. All Rights Reserved
              </div>
            </div>
          </div>
        </div>
      </div>
    </footer>
  )
}
