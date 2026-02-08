/** Basic conference display data */
export interface ConfMetadata {
  name: string;
  icon: string;
}

/** Basic paper display data */
export interface PaperMetadata {
  id: string;
  title: string;
}

/** Full paper data */
export interface Paper {
  title: string;
  authors: string[];
  abstract: string;
  pdf: string;
  poster?: string;
}
