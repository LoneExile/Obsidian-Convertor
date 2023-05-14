class ObsidianConvertor < Formula
  desc "A converter for Obsidian markdown files"
  homepage "https://github.com/LoneExile/Obsidian-Convertor"
  url "https://github.com/LoneExile/Obsidian-Convertor/releases/download/v0.1.0/oc"
  sha256 "3a266751e902b67d16ce7c4932fdcd363c49a48c7a21b7c6354cfbbcc45026b1"
  version "0.1.0"

  def install
    bin.install "oc"
  end

  test do
    system "#{bin}/oc", "--help"
  end
end
