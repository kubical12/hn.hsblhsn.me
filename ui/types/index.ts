export type FeedT = {
  feedItems: Array<FeedItemT | null>;
  totalPages: number;
  currentPage: number;
};

export type FeedItemT = {
  id: number;
  title: string;
  body: string;
  __html: string;
  link: string;
  images: ImageT[];
  totalPoints: number;
  totalComments: number;
};

export type ImageT = {
  url: string;
  width: number;
  height: number;
  alt: string;
  type: string;
};

export type ErrorT = {
  message: string;
  reason?: string;
};
